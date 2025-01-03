package interfaces

import (
	"context"
	"eden/modules/profile/domain"
	"eden/modules/profile/infrastructure/queue/message"
)

type ProfileService interface {
	CreateOrUpdateProfile(ctx context.Context, profile *domain.Profile) error
	GetProfileByID(ctx context.Context, id uint) (*domain.Profile, error)
}

type PhotoService interface {
	CreatePhoto(ctx context.Context, photo *domain.Photo) error
	GetPhotosByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error)
}

type StreamForgeMessageProcessor interface {
	Process(ctx context.Context, msg message.StreamForgeMessage) error
}
