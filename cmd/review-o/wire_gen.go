// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review-o/internal/biz"
	"review-o/internal/conf"
	"review-o/internal/data"
	"review-o/internal/server"
	"review-o/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, register *conf.Register, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(register)
	reviewClient := data.NewRpcClient(discovery)
	dataData, cleanup, err := data.NewData(confData, reviewClient, logger)
	if err != nil {
		return nil, nil, err
	}
	appealRepo := data.NewreAppealRepo(dataData, logger)
	reAppeal := biz.NewReAppeal(appealRepo, logger)
	appealService := service.NewAppealService(reAppeal)
	grpcServer := server.NewGRPCServer(confServer, appealService, logger)
	httpServer := server.NewHTTPServer(confServer, appealService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
