package integrationtest

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/alecsavvy/clockwise/sdk"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
)

var discprovUrls = []string{"http://0.0.0.0:6603"}

func randomDiscprov() string {
	randomIndex := rand.IntN(len(discprovUrls))
	return discprovUrls[randomIndex]
}

func newSdk() *sdk.ClockwiseSdk {
	node := randomDiscprov()
	sdk := sdk.NewSdk(fmt.Sprintf("%s/query", node))
	return sdk
}

var ks = keystore.NewKeyStore("./tmp/it/ks", keystore.StandardScryptN, keystore.StandardScryptP)

func generateWallet() (accounts.Account, *keystore.KeyStore) {
	password := uuid.New().String()
	account, err := ks.NewAccount(password)
	if err != nil {
			log.Fatalf("Failed to create new account: %v", err)
	}
	return account, ks
}
