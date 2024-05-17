package handlers

import (
	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// handles safe processing of cqrs commands and emits events on successful processing
// if a transaction gets here and fails it will be dropped instead of erroring
func RootHandler(qtx *db.Queries, createdAt int32, txs [][]byte) ([]*abcitypes.ExecTxResult, error) {
	var txResults = make([]*abcitypes.ExecTxResult, len(txs))

	if len(txs) <= 0 {
		return txResults, nil
	}

	// init tx results
	for i, tx := range txs {
		baseCmd, err := chainutils.FromTxBytes[commands.Command[any]](tx)
		if err != nil {
			return nil, utils.AppError("could not parse a tx as a command", err)
		}

		operation := baseCmd.Operation

		switch operation {
		case commands.Operation{Action: commands.CREATE, Entity: commands.USER}:
			txResult, err := HandleCreateUser(qtx, createdAt, tx)
			if err != nil {
				return nil, utils.AppError("cannot handle create user", err)
			}
			txResults[i] = txResult
		case commands.Operation{Action: commands.CREATE, Entity: commands.TRACK}:
			txResult, err := HandleCreateTrack(qtx, createdAt, tx)
			if err != nil {
				return nil, utils.AppError("cannot handle create track", err)
			}
			txResults[i] = txResult
		case commands.Operation{Action: commands.CREATE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.CREATE, Entity: commands.REPOST}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.REPOST}:
		default:
			return nil, utils.AppError("found transaction with invalid operation", nil)
		}
	}
	return txResults, nil
}
