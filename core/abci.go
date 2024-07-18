package core

import (
	"context"
	"errors"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

const (
	CodeTypeNotOK = 1
)

// Collects info about the chain node and returns it to other nodes.
func (c *Core) Info(context.Context, *abcitypes.InfoRequest) (*abcitypes.InfoResponse, error) {
	return &abcitypes.InfoResponse{}, nil
}

// Instructions for how the chain should initialize.
func (c *Core) InitChain(context.Context, *abcitypes.InitChainRequest) (*abcitypes.InitChainResponse, error) {
	return &abcitypes.InitChainResponse{}, nil
}

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (c *Core) CheckTx(ctx context.Context, req *abcitypes.CheckTxRequest) (*abcitypes.CheckTxResponse, error) {
	if err := c.validateTx(ctx, req.GetTx()); err != nil {
		c.logger.Error("error in check tx", "error", err)
		return &abcitypes.CheckTxResponse{Code: CodeTypeNotOK, Log: err.Error()}, nil
	}
	return &abcitypes.CheckTxResponse{Code: abcitypes.CodeTypeOK}, nil
}

// Prepares a new block proposal for the network
func (c *Core) PrepareProposal(_ context.Context, proposal *abcitypes.PrepareProposalRequest) (*abcitypes.PrepareProposalResponse, error) {
	// TODO: reorder transactions in here
	return &abcitypes.PrepareProposalResponse{Txs: proposal.Txs}, nil
}

// Processes block proposal from the network created by PrepareProposal
func (c *Core) ProcessProposal(_ context.Context, proposal *abcitypes.ProcessProposalRequest) (*abcitypes.ProcessProposalResponse, error) {
	return &abcitypes.ProcessProposalResponse{Status: abcitypes.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

// Prepares a block for commitment and provides final validation
func (c *Core) FinalizeBlock(ctx context.Context, rfb *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	txResults, err := c.indexTxs(ctx, rfb)
	if err != nil {
		c.logger.Errorf("CONSENSUS ERROR %s", err)
		return nil, err
	}
	return &abcitypes.FinalizeBlockResponse{TxResults: txResults}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (c *Core) Commit(ctx context.Context, req *abcitypes.CommitRequest) (*abcitypes.CommitResponse, error) {
	resp := &abcitypes.CommitResponse{}
	/**
	// TODO: check if indexer is up to date here, only prune once indexer is up to date.
	// i.e. a node can seed postgres after indexing
	if app.RetainBlocks > 0 && app.state.Height >= app.RetainBlocks {
		resp.RetainHeight = app.state.Height - app.RetainBlocks + 1
		}
	*/
	latestBlock, err := c.rpc.Block(ctx, nil)
	if err != nil {
		return nil, err
	}
	if c.RetainBlocks > 0 && latestBlock.Block.Height >= c.RetainBlocks {
		resp.RetainHeight = latestBlock.Block.Height - c.RetainBlocks + 1
	}

	if err := c.commitInProgressTx(ctx); err != nil {
		c.logger.Error("failure to commit tx", "error", err)
	}

	return resp, nil
}

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (c *Core) Query(_ context.Context, req *abcitypes.QueryRequest) (*abcitypes.QueryResponse, error) {
	return nil, errors.New("internal query not supported")
}

// ListSnapshots implements types.Application.
func (c *Core) ListSnapshots(context.Context, *abcitypes.ListSnapshotsRequest) (*abcitypes.ListSnapshotsResponse, error) {
	return &abcitypes.ListSnapshotsResponse{}, nil
}

// LoadSnapshotChunk implements types.Application.
func (c *Core) LoadSnapshotChunk(context.Context, *abcitypes.LoadSnapshotChunkRequest) (*abcitypes.LoadSnapshotChunkResponse, error) {
	return &abcitypes.LoadSnapshotChunkResponse{}, nil
}

// OfferSnapshot implements types.Application.
func (c *Core) OfferSnapshot(context.Context, *abcitypes.OfferSnapshotRequest) (*abcitypes.OfferSnapshotResponse, error) {
	return &abcitypes.OfferSnapshotResponse{}, nil
}

// ApplySnapshotChunk implements types.Application.
func (c *Core) ApplySnapshotChunk(context.Context, *abcitypes.ApplySnapshotChunkRequest) (*abcitypes.ApplySnapshotChunkResponse, error) {
	return &abcitypes.ApplySnapshotChunkResponse{}, nil
}

// ExtendVote implements types.Application.
func (c *Core) ExtendVote(context.Context, *abcitypes.ExtendVoteRequest) (*abcitypes.ExtendVoteResponse, error) {
	return &abcitypes.ExtendVoteResponse{}, nil
}

// VerifyVoteExtension implements types.Application.
func (c *Core) VerifyVoteExtension(_ context.Context, req *abcitypes.VerifyVoteExtensionRequest) (*abcitypes.VerifyVoteExtensionResponse, error) {
	return &abcitypes.VerifyVoteExtensionResponse{}, nil
}
