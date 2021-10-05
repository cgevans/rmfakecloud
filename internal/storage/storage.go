package storage

import (
	"errors"
	"io"
	"time"

	"github.com/ddvk/rmfakecloud/internal/messages"
	"github.com/ddvk/rmfakecloud/internal/model"
)

type ExportOption int

const DocumentType = "DocumentType"
const CollectionType = "CollectionType"
const MetadataFileExt = ".metadata"

const (
	ExportWithAnnotations ExportOption = iota
	ExportOnlyAnnotations
)

var ErrorNotFound = errors.New("not found")
var ErrorWrongGeneration = errors.New("wrong generation")

// DocumentStorer stores documents
type DocumentStorer interface {
	StoreDocument(uid, docid string, s io.ReadCloser) error
	RemoveDocument(uid, docid string) error
	GetDocument(uid, docid string) (io.ReadCloser, error)
	ExportDocument(uid, docid, outputType string, exportOption ExportOption) (io.ReadCloser, error)

	GetStorageURL(uid, docid, urltype string) (string, time.Time, error)
}

// BlobStorage stuff for sync15
type BlobStorage interface {
	GetBlobURL(uid, docid string) (string, time.Time, error)

	StoreBlob(uid, blobId string, s io.Reader, matchGeneration int64) (int64, error)
	LoadBlob(uid, blobId string) (io.ReadCloser, int64, error)
}

// MetadataStorer manages document metadata
type MetadataStorer interface {
	UpdateMetadata(uid string, r *messages.RawDocument) error
	GetAllMetadata(uid string) ([]*messages.RawDocument, error)
	GetMetadata(uid, docid string) (*messages.RawDocument, error)
}

// UserStorer holds informations about users
type UserStorer interface {
	GetUsers() ([]*model.User, error)
	GetUser(string) (*model.User, error)
	RegisterUser(u *model.User) error
	UpdateUser(u *model.User) error
}

// Document represents a document in storage
type Document struct {
	ID     string
	Type   string
	Parent string
	Name   string
}
