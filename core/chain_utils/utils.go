package chainutils

import (
	"encoding/json"

	"github.com/alecsavvy/clockwise/utils"
)

func ToTxBytes(tx interface{}) ([]byte, error) {
	txBytes, err := json.Marshal(tx)
	if err != nil {
		return nil, utils.AppError("could not marshal tx to json bytes", err)
	}
	return txBytes, nil
}

func FromTxBytes(jsonBytes []byte) (interface{}, error) {
	var result interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
