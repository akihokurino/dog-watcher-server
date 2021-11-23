package graph

import (
	"live-server/infra/firebase"
	"context"
)

const authUidStoreKey = "__auth_uid_store_key__"

type ContextProvider interface {
	WithAuthUID(ctx context.Context, uid firebase.UID) context.Context
	MustAuthUID(ctx context.Context) firebase.UID
}

type contextProvider struct {
}

func NewContextProvider() ContextProvider {
	return &contextProvider{}
}

func (u *contextProvider) WithAuthUID(ctx context.Context, uid firebase.UID) context.Context {
	return context.WithValue(ctx, authUidStoreKey, uid)
}

func (u *contextProvider) MustAuthUID(ctx context.Context) firebase.UID {
	v, ok := ctx.Value(authUidStoreKey).(firebase.UID)
	if !ok {
		panic("not found uid")
	}
	return v
}
