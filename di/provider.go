// +build wireinject

package di

import (
	"canvas-server/graph"
	"canvas-server/infra/agora"
	"canvas-server/infra/firebase"
	"os"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	firebase.NewClient,
	provideAgoraClient,
	graph.NewResolver,
	graph.NewServer,
	graph.NewContextProvider,
	graph.NewAuthenticate,
	graph.NewCROS,
)

func provideAgoraClient() agora.Client {
	return agora.NewClient(os.Getenv("AGORA_APP_ID"), os.Getenv("AGORA_CERT_ID"))
}

func ResolveGraphQL() graph.Server {
	panic(wire.Build(providerSet))
}
