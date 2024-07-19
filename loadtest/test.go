package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/alecsavvy/clockwise/sdk"
	"github.com/bxcodec/faker/v3"
)

// runs a test sequence
// creates three users
// uploads a track for two of those users
// users repost and follow each other
// uses random sdks between calls to land on different nodes between actions
func testSequence(stats *Stats) error {
	u1handle, u1addr, u1bio := generateTestUser()
	u2handle, u2addr, u2bio := generateTestUser()
	u3handle, u3addr, u3bio := generateTestUser()

	sdk := newSdk()

	_, err := sdk.CreateUser(u1handle, u1addr, u1bio)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	_, err = sdk.CreateUser(u2handle, u2addr, u2bio)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	_, err = sdk.CreateUser(u3handle, u3addr, u3bio)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	sdk = newSdk()

	t1title, t1su, t1u, t1desc := generateTestTrack(u1addr)
	t2title, t2su, t2u, t2desc := generateTestTrack(u2addr)

	track, err := sdk.CreateTrack(t1title, t1su, t1u, t1desc)
	if err != nil {
		return err
	}
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	_, err = sdk.RepostTrack(track.CreateTrack.Id, u3addr)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	track, err = sdk.CreateTrack(t2title, t2su, t2u, t2desc)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	_, err = sdk.RepostTrack(track.CreateTrack.Id, u3addr)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	// reset sdk to yet another different node
	sdk = newSdk()

	_, err = sdk.FollowUser(u3addr, u1addr)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	_, err = sdk.FollowUser(u3addr, u2addr)
	stats.recordStat(sdk.GetEndpoint(), err)
	if err != nil {
		return err
	}

	return nil
}

func newSdk() *sdk.ClockwiseSdk {
	node := randomDiscprov()
	sdk := sdk.NewSdk(fmt.Sprintf("%s/query", node))
	return sdk
}

// eventually return wallet when signing is required
func generateTestUser() (string, string, string) {
	account, _ := generateWallet()
	handle := faker.Username()
	address := account.Address.Hex()
	bio := faker.Sentence()
	return handle, address, bio
}

func generateTestTrack(userId string) (string, string, string, string) {
	title := fmt.Sprintf("%s' %s", userId, faker.Sentence())
	streamUrl := faker.DomainName()
	description := faker.Paragraph()
	return title, streamUrl, userId, description
}

var discprovUrls = []string{"http://node-0:26659", "http://node-1:26659", "http://node-2:26659", "http://node-3:26659", "http://node-4:26659", "http://node-5:26659", "http://node-6:26659"}

func randomDiscprov() string {
	randomIndex := rand.IntN(len(discprovUrls))
	return discprovUrls[randomIndex]
}
