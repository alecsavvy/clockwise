package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/alecsavvy/clockwise/sdk"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/google/uuid"
)

// sends random entity manager transations to random nodes
func main() {
	logger := utils.NewLogger(nil)
	for {
		sdk := sdk.NewSdk(fmt.Sprintf("%s/query", randomDiscprov()))

		requestId := uuid.NewString()
		userId := randomIntID()
		entityId := randomIntID()
		signer := uuid.NewString()
		entityType := randomEntity()
		action := randomAction()
		metadata := "metadata"

		_, err := sdk.ManageEntity(
			requestId,
			userId,
			signer,
			entityType,
			entityId,
			metadata,
			action,
		)

		if err != nil {
			logger.Error("error sending manage entity", "error", err)
		}
		time.Sleep(1 * time.Second)
	}
}

var discprovUrls = []string{"http://node-0:26659", "http://node-1:26659", "http://node-2:26659", "http://node-3:26659", "http://node-4:26659", "http://node-5:26659", "http://node-6:26659"}

func randomDiscprov() string {
	randomIndex := rand.IntN(len(discprovUrls))
	return discprovUrls[randomIndex]
}

var entities = []string{"User", "Track", "Playlist"}
var actions = []string{"Create", "Update", "Repost", "Follow", "Unfollow", "Unrepost", "Delete"}

func randomEntity() string {
	randomIndex := rand.IntN(len(entities))
	return entities[randomIndex]
}

func randomAction() string {
	randomIndex := rand.IntN(len(actions))
	return actions[randomIndex]
}

func randomIntID() int {
	return rand.Int()
}
