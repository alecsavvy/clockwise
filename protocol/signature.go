package protocol

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"

	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"
)

func Sign(pk *ecdsa.PrivateKey, msg proto.Message) (*gen.EnvelopeV2, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	hash32 := sha256.Sum256(data)
	hash := hash32[:]

	signature, err := crypto.Sign(hash, pk)
	if err != nil {
		return nil, err
	}

	return &gen.EnvelopeV2{
		Data:      data,
		Hash:      hash,
		Signature: signature,
	}, nil
}

func Recover[T proto.Message](msg *gen.EnvelopeV2, data T) (*common.Address, error) {
	pubKey, err := crypto.SigToPub(msg.Hash, msg.Signature)
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(*pubKey)

	dataHash := sha256.Sum256(msg.Data)
	if dataHash != [32]byte(msg.Hash) {
		return &address, errors.New("hashed data does not match, tampered with")
	}

	if err = proto.Unmarshal(msg.Data, data); err != nil {
		return &address, err
	}

	return &address, nil
}
