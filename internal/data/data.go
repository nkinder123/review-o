package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	v1 "review-o/api/review/v1"
	"review-o/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewRpcClient, NewData, NewreAppealRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rpcClinet v1.ReviewClient
	log       *log.Helper
}

// NewData .
func NewData(c *conf.Data, rpcClient v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rpcClinet: rpcClient, log: log.NewHelper(logger)}, cleanup, nil
}

// 获取rpc的client
func NewRpcClient(register registry.Discovery) v1.ReviewClient {
	insecure, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///review.service"),
		grpc.WithDiscovery(register),
		grpc.WithTimeout(3600*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(insecure)
}

// 获取consul的discovery
func NewDiscovery(c *conf.Register) registry.Discovery {
	config := api.DefaultConfig()
	config.Address = c.Consul.Address
	config.Scheme = c.Consul.Scheme
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	return consul.New(client)

}
