// +build wireinject

package di

import (
	"canvas-server/graph"
	"canvas-server/infra/firebase"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	firebase.NewClient,
	graph.NewResolver,
	graph.NewServer,
	graph.NewContextProvider,
	graph.NewAuthenticate,
	graph.NewCROS,
)

func ResolveGraphQL() graph.Server {
	panic(wire.Build(providerSet))
}
