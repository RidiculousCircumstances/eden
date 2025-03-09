package interfaces

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

type StorageClient interface {
	UploadObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, opts interface{}) error
	DeleteObject(ctx context.Context, bucketName, objectName string) error
	GetObject(ctx context.Context, bucketName, objectName string, opts interface{}) (*minio.Object, error)
}
