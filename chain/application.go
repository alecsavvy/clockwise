package chain

import (
	"context"

	"github.com/cometbft/cometbft/abci/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Clockwise implementation of the abci interface for cometbft
// everything related to consensus MUST go through here
// https://docs.cometbft.com/v0.38/spec/abci/abci++_methods
type Application struct {
}

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (a *Application) CheckTx(context.Context, *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	panic("check if valid entity, etc")
}

// Writes the state changes to the database after checking and finalizing a block
func (a *Application) Commit(context.Context, *types.RequestCommit) (*types.ResponseCommit, error) {
	panic("commit transaction to postgres")
}

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(context.Context, *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	panic("unimplemented")
}

// Prepares a new block proposal for the network
func (a *Application) PrepareProposal(_ context.Context, proposal *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	return &abcitypes.ResponsePrepareProposal{Txs: proposal.Txs}, nil
}

// Processes block proposal from the network created by PrepareProposal
func (a *Application) ProcessProposal(context.Context, *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	return &abcitypes.ResponseProcessProposal{Status: abcitypes.ResponseProcessProposal_ACCEPT}, nil
}

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (a *Application) Query(_ context.Context, req *types.RequestQuery) (*types.ResponseQuery, error) {
	resp := abcitypes.ResponseQuery{Key: req.Data}
	return &resp, nil
}

// ExtendVote implements types.Application.
func (a *Application) ExtendVote(context.Context, *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	return &abcitypes.ResponseExtendVote{}, nil
}

// VerifyVoteExtension implements types.Application.
func (a *Application) VerifyVoteExtension(_ context.Context, req *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	return &abcitypes.ResponseVerifyVoteExtension{}, nil
}

// Info implements types.Application.
func (a *Application) Info(context.Context, *types.RequestInfo) (*types.ResponseInfo, error) {
	return &abcitypes.ResponseInfo{}, nil
}

// InitChain implements types.Application.
func (a *Application) InitChain(context.Context, *types.RequestInitChain) (*types.ResponseInitChain, error) {
	return &abcitypes.ResponseInitChain{}, nil
}

// ListSnapshots implements types.Application.
func (a *Application) ListSnapshots(context.Context, *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	return &abcitypes.ResponseListSnapshots{}, nil
}

// LoadSnapshotChunk implements types.Application.
func (a *Application) LoadSnapshotChunk(context.Context, *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	return &abcitypes.ResponseLoadSnapshotChunk{}, nil
}

// OfferSnapshot implements types.Application.
func (a *Application) OfferSnapshot(context.Context, *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	return &abcitypes.ResponseOfferSnapshot{}, nil
}

// ApplySnapshotChunk implements types.Application.
func (a *Application) ApplySnapshotChunk(context.Context, *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	return &abcitypes.ResponseApplySnapshotChunk{}, nil
}

// compile time check for abci compatibility
var _ abcitypes.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{}
}
