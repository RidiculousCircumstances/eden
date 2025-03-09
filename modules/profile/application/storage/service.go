package storage

import (
	"context"
	storageIntf "eden/modules/profile/application/storage/interfaces"
	"github.com/minio/minio-go/v7"
	"io"
)

type Service struct {
	client storageIntf.StorageClient
}

func NewService(client storageIntf.StorageClient) *Service {
	return &Service{client: client}
}

func (s *Service) UploadObject(ctx context.Context, bucketName string, objectName string, size int64, reader io.Reader, opts minio.PutObjectOptions) error {
	return s.client.UploadObject(ctx, bucketName, objectName, reader, size, opts)
}
