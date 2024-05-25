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

func (sdk *ClockwiseSdk) ManageEntity(requestId string,
	userId int,
	signer string,
	entityType string,
	entityId int,
	metadata string,
	action string) (*gqlclient.ManageEntityResponse, error) {
	return gqlclient.ManageEntity(context.Background(), *sdk.client, requestId, userId, signer, entityType, entityId, metadata, action)
}
