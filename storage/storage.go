package storage

import (
	"io"

	"github.com/alecsavvy/clockwise/types"
)

type StorageService interface {
	UploadFile(r io.Reader) (types.CID, error)
	GetFile(types.CID) (io.ReadCloser, error)
	GetFileMetadata(types.CID) (interface{}, error)
	UpdateFile(cid types.CID, metadata interface{}, r io.Reader) (types.CID, error)
	StreamFile(cid types.CID, w io.Writer) error
}
