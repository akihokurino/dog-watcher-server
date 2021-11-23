package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"live-server/graph/generated"
	"live-server/graph/model"
)

func (r *queryResolver) RtcToken(ctx context.Context) (*model.RTCToken, error) {
	uid := r.contextProvider.MustAuthUID(ctx)
	token, err := r.agoraClient.GetRTCToken(uid)
	if err != nil {
		return nil, err
	}

	return &model.RTCToken{
		Token: token,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
