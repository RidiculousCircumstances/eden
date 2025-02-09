package interfaces

import (
	"context"
	"eden/modules/profile/domain"
)

type PhotoRepository interface {
	// Create создает новую запись Photo.
	Create(ctx context.Context, photo *domain.Photo) error

	// GetByID возвращает Photo по его ID.
	GetByID(ctx context.Context, id uint) (*domain.Photo, error)

	ExistsByIndexID(ctx context.Context, id uint) (bool, error)

	GetIDByIndexID(ctx context.Context, id uint32) (uint, error)

	// GetByProfileID возвращает все Photos, связанные с указанным ProfileID.
	GetByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error)

	GetProfileByPhotoIndexID(ctx context.Context, indexID uint32) (*domain.Profile, error)

	GetProfilesByPhotoIndexIDs(ctx context.Context, indexIDs []uint32, limit int) ([]domain.Profile, error)

	// Update обновляет данные существующего Photo.
	Update(ctx context.Context, photo *domain.Photo) error

	// Delete удаляет Photo по его ID.
	Delete(ctx context.Context, id uint) error
}

type ProfileRepository interface {
	// Create создает новый профиль.
	Create(ctx context.Context, profile *domain.Profile) error

	// GetByID возвращает профиль по его ID.
	GetByID(ctx context.Context, id uint) (*domain.Profile, error)

	// Update обновляет данные существующего профиля.
	Update(ctx context.Context, profile *domain.Profile) error

	// Delete удаляет профиль по его ID.
	Delete(ctx context.Context, id uint) error
}

type FaceRepository interface {
	Create(ctx context.Context, face *domain.Face) error
}
