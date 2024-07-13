//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/kabuto412rock/microblog/internal"
	"github.com/kabuto412rock/microblog/internal/config"
	"github.com/kabuto412rock/microblog/internal/controller"
)

func InitalizeServer() *internal.Server {

	wire.Build(config.ReadConfig, controller.NewEnv, internal.NewServer)
	return &internal.Server{}
}
