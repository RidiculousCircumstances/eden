package repository

import (
	"context"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
	"errors"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

// NewPhotoRepository создает новый экземпляр репозитория Photo.
func NewPhotoRepository(db *gorm.DB) interfaces.PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) Create(ctx context.Context, photo *domain.Photo) error {
	// Создает новую запись Photo в базе данных
	return r.db.WithContext(ctx).Create(photo).Error
}

func (r *photoRepository) GetByID(ctx context.Context, id uint) (*domain.Photo, error) {
	var photo domain.Photo
	err := r.db.WithContext(ctx).First(&photo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Если фото не найдено, возвращаем nil
	}
	return &photo, err
}

func (r *photoRepository) GetByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Where("profile_id = ?", profileID).Find(&photos).Error
	return photos, err
}

// В репозитории можно создать метод для выполнения запроса с JOIN.
func (r *photoRepository) GetProfileByPhotoIndexID(ctx context.Context, indexID uint) (*domain.Profile, error) {
	var profile domain.Profile
	err := r.db.WithContext(ctx).
		Joins("JOIN profiles ON profiles.id = photos.profile_id").
		Where("photos.index_id = ?", indexID).
		First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *photoRepository) Update(ctx context.Context, photo *domain.Photo) error {
	// Обновляет запись Photo в базе данных
	return r.db.WithContext(ctx).Save(photo).Error
}

func (r *photoRepository) Delete(ctx context.Context, id uint) error {
	// Удаляет запись Photo по его ID
	return r.db.WithContext(ctx).Delete(&domain.Photo{}, id).Error
}
