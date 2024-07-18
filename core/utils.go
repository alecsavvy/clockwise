package core

import (
	"github.com/cometbft/cometbft/libs/bytes"
	"github.com/cometbft/cometbft/types"
	"google.golang.org/protobuf/proto"
)

func ToTxHash(msg proto.Message) (string, error) {
	b, err := proto.Marshal(msg)
	if err != nil {
		return "", err
	}

	tx := types.Tx(b)
	hash := tx.Hash()
	hexBytes := bytes.HexBytes(hash)
	hashStr := hexBytes.String()
	
	return hashStr, nil
}
