package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"live-server/graph/generated"
	"live-server/graph/model"
)

func (r *mutationResolver) CreateAgoraToken(ctx context.Context, input model.CreateAgoraTokenInput) (*model.AgoraToken, error) {
	uid := r.contextProvider.MustAuthUID(ctx)
	token, err := r.agoraClient.GetRTCToken(uid, input.ChannelName, input.Role)
	if err != nil {
		return nil, err
	}

	return &model.AgoraToken{
		Token: token,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
