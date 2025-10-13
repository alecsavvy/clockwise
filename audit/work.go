package audit

type ProofOfWorkService interface {
	CreatePOWRollup() error
	GetPOWRollup(rollupID int64) error
	GetPOWRollups(address string, perPage int64, page int64) error
}
