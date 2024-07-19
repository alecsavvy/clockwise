package sdk

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/alecsavvy/clockwise/sdk/gqlclient"
)

type ClockwiseSdk struct {
	endpoint string
	client *graphql.Client
}

func NewSdk(gqlEndpoint string) *ClockwiseSdk {
	client := graphql.NewClient(gqlEndpoint, http.DefaultClient)
	return &ClockwiseSdk{
		endpoint: gqlEndpoint,
		client: &client,
	}
}

func (sdk *ClockwiseSdk) GetEndpoint() string {
	return sdk.endpoint
}

func (sdk *ClockwiseSdk) CreateUser(handle, address, bio string) (*gqlclient.CreateUserResponse, error) {
	return gqlclient.CreateUser(context.Background(), *sdk.client, handle, address, bio)
}

func (sdk *ClockwiseSdk) CreateTrack(title, streamURL, userID, description string) (*gqlclient.CreateTrackResponse, error) {
	return gqlclient.CreateTrack(context.Background(), *sdk.client, title, streamURL, userID, description)
}

func (sdk *ClockwiseSdk) FollowUser(followerID, followeeID string) (*gqlclient.FollowResponse, error) {
	return gqlclient.Follow(context.Background(), *sdk.client, followerID, followeeID)
}

func (sdk *ClockwiseSdk) RepostTrack(trackID, reposterID string) (*gqlclient.RepostResponse, error) {
	return gqlclient.Repost(context.Background(), *sdk.client, trackID, reposterID)
}
