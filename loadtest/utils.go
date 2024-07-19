package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
)

var ks = keystore.NewKeyStore("/tmp/lt/keystore", keystore.StandardScryptN, keystore.StandardScryptP)

func generateWallet() (accounts.Account, *keystore.KeyStore) {
	password := uuid.New().String()
	account, err := ks.NewAccount(password)
	if err != nil {
			log.Fatalf("Failed to create new account: %v", err)
	}
	return account, ks
}
