package graph

import (
	"live-server/infra/agora"
	"live-server/infra/firebase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	contextProvider ContextProvider
	fireClient      firebase.Client
	agoraClient     agora.Client
}

func NewResolver(
	contextProvider ContextProvider,
	fireClient firebase.Client,
	agoraClient agora.Client) *Resolver {
	return &Resolver{
		contextProvider: contextProvider,
		fireClient:      fireClient,
		agoraClient:     agoraClient,
	}
}
