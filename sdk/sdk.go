package sdk

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/alecsavvy/clockwise/sdk/gqlclient"
)

type ClockwiseSdk struct {
	client *graphql.Client
}

func NewSdk(gqlEndpoint string) *ClockwiseSdk {
	client := graphql.NewClient(gqlEndpoint, http.DefaultClient)
	return &ClockwiseSdk{
		client: &client,
	}
}

func (sdk *ClockwiseSdk) CreateUser(handle, address, bio string) (*gqlclient.CreateUserResponse, error) {
	return gqlclient.CreateUser(context.Background(), *sdk.client, handle, address, bio)
}
