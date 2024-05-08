package chain

import (
	"context"

	"github.com/cometbft/cometbft/abci/types"
)

// Clockwise implementation of the abci interface for cometbft
// everything related to consensus MUST go through here
// https://docs.cometbft.com/v0.38/spec/abci/abci++_methods
type Application struct {
}

// ApplySnapshotChunk implements types.Application.
func (a *Application) ApplySnapshotChunk(context.Context, *types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	panic("unimplemented")
}

// CheckTx implements types.Application.
func (a *Application) CheckTx(context.Context, *types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	panic("unimplemented")
}

// Commit implements types.Application.
func (a *Application) Commit(context.Context, *types.RequestCommit) (*types.ResponseCommit, error) {
	panic("unimplemented")
}

// ExtendVote implements types.Application.
func (a *Application) ExtendVote(context.Context, *types.RequestExtendVote) (*types.ResponseExtendVote, error) {
	panic("unimplemented")
}

// FinalizeBlock implements types.Application.
func (a *Application) FinalizeBlock(context.Context, *types.RequestFinalizeBlock) (*types.ResponseFinalizeBlock, error) {
	panic("unimplemented")
}

// Info implements types.Application.
func (a *Application) Info(context.Context, *types.RequestInfo) (*types.ResponseInfo, error) {
	panic("unimplemented")
}

// InitChain implements types.Application.
func (a *Application) InitChain(context.Context, *types.RequestInitChain) (*types.ResponseInitChain, error) {
	panic("unimplemented")
}

// ListSnapshots implements types.Application.
func (a *Application) ListSnapshots(context.Context, *types.RequestListSnapshots) (*types.ResponseListSnapshots, error) {
	panic("unimplemented")
}

// LoadSnapshotChunk implements types.Application.
func (a *Application) LoadSnapshotChunk(context.Context, *types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	panic("unimplemented")
}

// OfferSnapshot implements types.Application.
func (a *Application) OfferSnapshot(context.Context, *types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error) {
	panic("unimplemented")
}

// PrepareProposal implements types.Application.
func (a *Application) PrepareProposal(context.Context, *types.RequestPrepareProposal) (*types.ResponsePrepareProposal, error) {
	panic("unimplemented")
}

// ProcessProposal implements types.Application.
func (a *Application) ProcessProposal(context.Context, *types.RequestProcessProposal) (*types.ResponseProcessProposal, error) {
	panic("unimplemented")
}

// Query implements types.Application.
func (a *Application) Query(context.Context, *types.RequestQuery) (*types.ResponseQuery, error) {
	panic("unimplemented")
}

// VerifyVoteExtension implements types.Application.
func (a *Application) VerifyVoteExtension(context.Context, *types.RequestVerifyVoteExtension) (*types.ResponseVerifyVoteExtension, error) {
	panic("unimplemented")
}

// compile time check for abci compatibility
// var _ abcitypes.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{}
}
