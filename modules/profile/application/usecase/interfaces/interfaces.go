package interfaces

import (
	"context"
	"eden/modules/profile/domain"
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
