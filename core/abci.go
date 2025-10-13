package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

var _ abcitypes.Application = (*Core)(nil)

// ApplySnapshotChunk implements types.Application.
func (c *Core) ApplySnapshotChunk(ctx context.Context, req *abcitypes.ApplySnapshotChunkRequest) (*abcitypes.ApplySnapshotChunkResponse, error) {
	panic("unimplemented")
}

// CheckTx implements types.Application.
func (c *Core) CheckTx(ctx context.Context, req *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
	panic("unimplemented")
}

// Commit implements types.Application.
func (c *Core) Commit(ctx context.Context, req *abcitypes.CommitRequest) (*abcitypes.CommitResponse, error) {
	panic("unimplemented")
}

// ExtendVote implements types.Application.
func (c *Core) ExtendVote(ctx context.Context, req *abcitypes.ExtendVoteRequest) (*abcitypes.ExtendVoteResponse, error) {
	panic("unimplemented")
}

// FinalizeBlock implements types.Application.
func (c *Core) FinalizeBlock(ctx context.Context, req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	panic("unimplemented")
}

// Info implements types.Application.
func (c *Core) Info(ctx context.Context, req *abcitypes.InfoRequest) (*abcitypes.InfoResponse, error) {
	panic("unimplemented")
}

// InitChain implements types.Application.
func (c *Core) InitChain(ctx context.Context, req *abcitypes.InitChainRequest) (*abcitypes.InitChainResponse, error) {
	panic("unimplemented")
}

// ListSnapshots implements types.Application.
func (c *Core) ListSnapshots(ctx context.Context, req *abcitypes.ListSnapshotsRequest) (*abcitypes.ListSnapshotsResponse, error) {
	panic("unimplemented")
}

// LoadSnapshotChunk implements types.Application.
func (c *Core) LoadSnapshotChunk(ctx context.Context, req *abcitypes.LoadSnapshotChunkRequest) (*abcitypes.LoadSnapshotChunkResponse, error) {
	panic("unimplemented")
}

// OfferSnapshot implements types.Application.
func (c *Core) OfferSnapshot(ctx context.Context, req *abcitypes.OfferSnapshotRequest) (*abcitypes.OfferSnapshotResponse, error) {
	panic("unimplemented")
}

// PrepareProposal implements types.Application.
func (c *Core) PrepareProposal(ctx context.Context, req *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	panic("unimplemented")
}

// ProcessProposal implements types.Application.
func (c *Core) ProcessProposal(ctx context.Context, req *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
	panic("unimplemented")
}

// Query implements types.Application.
func (c *Core) Query(ctx context.Context, req *abcitypes.QueryRequest) (*abcitypes.QueryResponse, error) {
	panic("unimplemented")
}

// VerifyVoteExtension implements types.Application.
func (c *Core) VerifyVoteExtension(ctx context.Context, req *abcitypes.VerifyVoteExtensionRequest) (*abcitypes.VerifyVoteExtensionResponse, error) {
	panic("unimplemented")
}
