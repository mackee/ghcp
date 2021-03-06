//+build wireinject

// Package di provides dependency injection.
package di

import (
	"github.com/google/wire"
	"github.com/int128/ghcp/adaptors"
	"github.com/int128/ghcp/adaptors/cmd"
	"github.com/int128/ghcp/adaptors/env"
	githubAdaptor "github.com/int128/ghcp/adaptors/github"
	"github.com/int128/ghcp/adaptors/logger"
	githubInfrastructure "github.com/int128/ghcp/infrastructure/github"
	"github.com/int128/ghcp/usecases/btc"
	"github.com/int128/ghcp/usecases/commit"
	"github.com/int128/ghcp/usecases/fork"
)

func NewCmd() adaptors.Cmd {
	wire.Build(
		cmd.Set,
		env.Set,
		githubAdaptor.Set,
		logger.Set,

		githubInfrastructure.Set,

		btc.Set,
		commit.Set,
		fork.Set,
	)
	return nil
}
