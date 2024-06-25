//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kabuto412rock/microblog/config"
	"github.com/kabuto412rock/microblog/controller"
	"github.com/kabuto412rock/microblog/internal"
)

func InitalizeServer() *internal.Server {
	wire.Build(config.ReadConfig, controller.NewEnv, internal.NewServer)
	return &internal.Server{}
}
