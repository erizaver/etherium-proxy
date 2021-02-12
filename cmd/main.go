package main

import (
	"context"

	"github.com/golang/glog"

	"github.com/erizaver/etherium_proxy/internal/pkg/app"
)

func main() {
	ctx := context.Background()

	application := app.NewApp()

	if err := application.Run(ctx); err != nil {
		glog.Fatal("error running application", err)
	}
}
