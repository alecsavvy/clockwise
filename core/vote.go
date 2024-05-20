/*
vote.go

Contains consensus voting logic.
*/
package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// ExtendVote implements types.Application.
func (c *Core) ExtendVote(context.Context, *abcitypes.RequestExtendVote) (*abcitypes.ResponseExtendVote, error) {
	return &abcitypes.ResponseExtendVote{}, nil
}

// VerifyVoteExtension implements types.Application.
func (c *Core) VerifyVoteExtension(_ context.Context, req *abcitypes.RequestVerifyVoteExtension) (*abcitypes.ResponseVerifyVoteExtension, error) {
	return &abcitypes.ResponseVerifyVoteExtension{}, nil
}
