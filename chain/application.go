package chain

import (
	"context"

	"github.com/alecsavvy/clockwise/db"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Clockwise implementation of the abci interface for cometbft
// everything related to consensus MUST go through here
// https://docs.cometbft.com/v0.38/spec/abci/abci++_methods
type Application struct {
	queries   *db.Queries
	pool      *pgxpool.Pool
	currentTx pgx.Tx
}

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (a *Application) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	_ = req.GetTx()

	return &abcitypes.ResponseCheckTx{}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (a *Application) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	// check transactions again
	a.currentTx.Commit(ctx)
	return &abcitypes.ResponseCommit{}, nil
}

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	dbTx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer dbTx.Rollback(ctx)

	qtx := a.queries.WithTx(dbTx)

	// do db logic with qtx
	blockTime := pgtype.Date{
		Time:  rfb.Time,
		Valid: true,
	}
	qtx.InsertBlock(ctx, db.InsertBlockParams{
		Blocknumber: rfb.Height,
		Blocktime:   blockTime,
	})

	// set transaction to be committed in commit step
	a.currentTx = dbTx
	return &abcitypes.ResponseFinalizeBlock{}, nil
}

// Prepares a new block proposal for the network
func (a *Application) PrepareProposal(_ context.Context, proposal *abcitypes.RequestPrepareProposal) (*abcitypes.ResponsePrepareProposal, error) {
	return &abcitypes.ResponsePrepareProposal{Txs: proposal.Txs}, nil
}

// Processes block proposal from the network created by PrepareProposal
func (a *Application) ProcessProposal(context.Context, *abcitypes.RequestProcessProposal) (*abcitypes.ResponseProcessProposal, error) {
	return &abcitypes.ResponseProcessProposal{Status: abcitypes.ResponseProcessProposal_ACCEPT}, nil
}

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (a *Application) Query(_ context.Context, req *abcitypes.RequestQuery) (*abcitypes.ResponseQuery, error) {
	resp := abcitypes.ResponseQuery{Key: req.Data}
	return &resp, nil
}

// ExtendVote implements types.Application.
func (a *Application) ExtendVote(context.Context, *abcitypes.RequestExtendVote) (*abcitypes.ResponseExtendVote, error) {
	return &abcitypes.ResponseExtendVote{}, nil
}

// VerifyVoteExtension implements types.Application.
func (a *Application) VerifyVoteExtension(_ context.Context, req *abcitypes.RequestVerifyVoteExtension) (*abcitypes.ResponseVerifyVoteExtension, error) {
	return &abcitypes.ResponseVerifyVoteExtension{}, nil
}

// Info implements types.Application.
func (a *Application) Info(context.Context, *abcitypes.RequestInfo) (*abcitypes.ResponseInfo, error) {
	return &abcitypes.ResponseInfo{}, nil
}

// InitChain implements types.Application.
func (a *Application) InitChain(context.Context, *abcitypes.RequestInitChain) (*abcitypes.ResponseInitChain, error) {
	return &abcitypes.ResponseInitChain{}, nil
}

// ListSnapshots implements types.Application.
func (a *Application) ListSnapshots(context.Context, *abcitypes.RequestListSnapshots) (*abcitypes.ResponseListSnapshots, error) {
	return &abcitypes.ResponseListSnapshots{}, nil
}

// LoadSnapshotChunk implements types.Application.
func (a *Application) LoadSnapshotChunk(context.Context, *abcitypes.RequestLoadSnapshotChunk) (*abcitypes.ResponseLoadSnapshotChunk, error) {
	return &abcitypes.ResponseLoadSnapshotChunk{}, nil
}

// OfferSnapshot implements types.Application.
func (a *Application) OfferSnapshot(context.Context, *abcitypes.RequestOfferSnapshot) (*abcitypes.ResponseOfferSnapshot, error) {
	return &abcitypes.ResponseOfferSnapshot{}, nil
}

// ApplySnapshotChunk implements types.Application.
func (a *Application) ApplySnapshotChunk(context.Context, *abcitypes.RequestApplySnapshotChunk) (*abcitypes.ResponseApplySnapshotChunk, error) {
	return &abcitypes.ResponseApplySnapshotChunk{}, nil
}

// compile time check for abci compatibility
var _ abcitypes.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{}
}
