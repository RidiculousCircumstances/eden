package storage

import (
	"context"
	storageIntf "eden/modules/profile/application/storage/interfaces"
	"io"
)

type Service struct {
	client storageIntf.StorageClient
}

func NewService(client storageIntf.StorageClient) *Service {
	return &Service{client: client}
}

func (s *Service) UploadObject(ctx context.Context, bucketName string, objectName string, size int64, reader io.Reader) error {
	return s.client.UploadObject(ctx, bucketName, objectName, reader, size, nil)
}
