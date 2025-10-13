package types

import (
	"fmt"

	ipfscid "github.com/ipfs/go-cid"
)

type CID string

func StringToCID(s string) (CID, error) {
	_, err := ipfscid.Decode(s)
	if err != nil {
		return "", fmt.Errorf("invalid CID: %w", err)
	}
	return CID(s), nil
}
