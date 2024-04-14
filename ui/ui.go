package ui

type UI struct{}

func New() (*UI, error) {
	return &UI{}, nil
}
