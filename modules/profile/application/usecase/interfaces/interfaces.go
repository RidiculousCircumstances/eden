package interfaces

import (
	"context"
	"eden/modules/profile/domain"
	"github.com/minio/minio-go/v7"
	"io"
)

type ProfileService interface {
	CreateOrUpdateProfile(ctx context.Context, profile *domain.Profile) error
	GetProfileByID(ctx context.Context, id uint) (*domain.Profile, error)
}

type PhotoService interface {
	CreatePhoto(ctx context.Context, photo *domain.Photo) error
	GetPhotoIdByIndexId(ctx context.Context, indexId uint32) (uint, error)
	GetPhotosByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error)
	GetProfileByIndexId(ctx context.Context, indexId uint32) (*domain.Profile, error)
	GetProfilesByIndexIds(ctx context.Context, indexIds []uint32) ([]domain.Profile, error)
}

type FaceService interface {
	CreateFace(ctx context.Context, face *domain.Face) error
}

type TakeSnapshot interface {
	Process(ctx context.Context) (string, error)
}

type StorageService interface {
	UploadObject(ctx context.Context, bucketName string, objectName string, size int64, reader io.Reader, opts minio.PutObjectOptions) error
}
