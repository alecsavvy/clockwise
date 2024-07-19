package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
)

func generateWallet() (accounts.Account, *keystore.KeyStore) {
	ks := keystore.NewKeyStore("./tmp/lt/ks", keystore.StandardScryptN, keystore.StandardScryptP)
	password := uuid.New().String()
	account, err := ks.NewAccount(password)
	if err != nil {
			log.Fatalf("Failed to create new account: %v", err)
	}
	return account, ks
}
