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
	"github.com/erizaver/etherium_proxy/internal/app/ethservice"
	"github.com/erizaver/etherium_proxy/internal/pkg/ethcloudflareclient"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

const (
	timeout               = 5 * time.Second
	getBlockCloudflareUrl = "https://cloudflare-eth.com"
	cacheSize             = 128
)

type Application struct {
	//Facades
	EthApi *etherium.EthApi

	//Services
	EthService *ethservice.EthService

	//Clients
	EthCloudflareClient *ethcloudflareclient.EthCloudflareClient

	Mux        *runtime.ServeMux
	GrpcServer *grpc.Server
}

func NewApp() *Application {
	app := &Application{}

	httpCli := &http.Client{Timeout: timeout}
	cache, err := lru.New(cacheSize)
	if err != nil {
		glog.Fatal("unable to start cache", err)
	}

	app.EthCloudflareClient = ethcloudflareclient.NewEthCloudflareClient(httpCli, getBlockCloudflareUrl)
	app.EthService = ethservice.NewEthService(app.EthCloudflareClient, cache)
	app.EthApi = etherium.NewEthApi(app.EthService)

	app.Mux = runtime.NewServeMux()
	app.GrpcServer = grpc.NewServer()

	return app
}

func (a *Application) Run(ctx context.Context) error {
	err := api.RegisterEthServiceHandlerServer(ctx, a.Mux, a.EthApi)
	if err != nil {
		return errors.Wrap(err, "unable to start server")
	}

	api.RegisterEthServiceServer(a.GrpcServer, a.EthApi)

	//not sure if we need this, since i have no information about this API usage and RPS, only need this if RPS super low
	//also we might want to use it once, since we need to get last block ID
	a.EthApi.WarmUpLatestBlockNumber()

	return http.ListenAndServe(":8080", a.Mux)
}
