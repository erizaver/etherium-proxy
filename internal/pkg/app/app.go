package app

import (
	"context"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/erizaver/etherium_proxy/internal/app/etherium"
	"github.com/erizaver/etherium_proxy/internal/pkg/ethcloudflareclient"
	"github.com/erizaver/etherium_proxy/internal/pkg/ethservice"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

//we can move this to config, but i was running low on time :)
const (
	timeout               = 5 * time.Second
	getBlockCloudflareUrl = "https://cloudflare-eth.com"
	cacheSize             = 128
)

type Application struct {
	EthFacade *etherium.EthFacade

	EthService *ethservice.EthService
}

func NewApp() *Application {
	app := &Application{}

	httpCli := &http.Client{Timeout: timeout}

	cache, err := lru.New(cacheSize)
	if err != nil {
		glog.Fatal("unable to start user server", err)
	}
	ethCloudflareClient := ethcloudflareclient.NewEthCloudflareClient(httpCli, getBlockCloudflareUrl)

	ethService := ethservice.NewEthService(ethCloudflareClient, cache)


	app.EthService = ethService

	ethFacade := etherium.NewEthFacade(app.EthService)
	app.EthFacade = ethFacade

	//not sure if we need this, since i have no information about this API usage and RPS, only need this if RPS super low
	//also we might want to use it once, since we need to get last block ID
	ethFacade.WarmUpLatestBlockNumber()

	return app
}

func (a *Application) Run(ctx context.Context) error {
	mux := runtime.NewServeMux()
	grpcSrv := grpc.NewServer()

	err := api.RegisterEthServiceHandlerServer(ctx, mux, a.EthFacade)
	if err != nil {
		return errors.Wrap(err, "unable to start user server")
	}

	api.RegisterEthServiceServer(grpcSrv, a.EthFacade)

	return http.ListenAndServe(":8080", mux)
}
