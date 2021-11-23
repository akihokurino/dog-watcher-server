package graph

import (
	"canvas-server/infra/firebase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	contextProvider ContextProvider
	fireClient      firebase.Client
}

func NewResolver(
	contextProvider ContextProvider,
	fireClient firebase.Client) *Resolver {
	return &Resolver{
		contextProvider: contextProvider,
		fireClient:      fireClient,
	}
}
