//go:build wireinject

package main

import (
	"github.com/MuxiKeStack/be-feed/events"
	"github.com/MuxiKeStack/be-feed/grpc"
	"github.com/MuxiKeStack/be-feed/ioc"
	"github.com/MuxiKeStack/be-feed/repository"
	"github.com/MuxiKeStack/be-feed/repository/dao"
	"github.com/MuxiKeStack/be-feed/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		// grpc
		ioc.InitGRPCxKratosServer,
		grpc.NewFeedServiceServer,
		service.NewFeedService,
		repository.NewFeedEventRepository,
		dao.NewGORMFeedPushEventDAO,
		// consumer
		ioc.InitConsumers,
		events.NewFeedEventConsumer,
		// 第三方
		ioc.InitDB,
		ioc.InitEtcdClient,
		ioc.InitLogger,
		ioc.InitKafka,
	)
	return &App{}
}
