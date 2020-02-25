//+build wireinject

package main

import (
	"server2/cmd/handler"
	"server2/internal/storages"
	"server2/internal/token"

	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	storages.NewService,
)

var controllerSet = wire.NewSet(
	token.NewController,
	storages.NewController,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
)

var appSet = wire.NewSet(
	serviceSet,
	controllerSet,
	handlerSet,
)

func initHandler() *handler.Handler {
	wire.Build(appSet)
	return &handler.Handler{}
}
