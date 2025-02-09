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

func (r *photoRepository) ExistsByIndexID(ctx context.Context, id uint) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.Photo{}).Where("index_id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *photoRepository) GetIDByIndexID(ctx context.Context, indexID uint32) (uint, error) {
	var id uint
	err := r.db.WithContext(ctx).
		Model(&domain.Photo{}).
		Where("index_id = ?", indexID).
		Pluck("id", &id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil // Если запись не найдена, возвращаем 0
	}
	if err != nil {
		return 0, err // Возвращаем ошибку в случае других проблем
	}
	return id, nil
}

func (r *photoRepository) GetByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.WithContext(ctx).Where("profile_id = ?", profileID).Find(&photos).Error
	return photos, err
}

// В репозитории можно создать метод для выполнения запроса с JOIN.
func (r *photoRepository) GetProfileByPhotoIndexID(ctx context.Context, indexID uint32) (*domain.Profile, error) {
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

func (r *photoRepository) GetProfilesByPhotoIndexIDs(ctx context.Context, indexIDs []uint32, limit int) ([]domain.Profile, error) {
	if len(indexIDs) == 0 {
		return nil, nil // Если список пуст, возвращаем nil
	}

	var profiles []domain.Profile

	err := r.db.WithContext(ctx).
		Distinct(
			"profiles.id",
			"profiles.source_id",
			"profiles.city_id",
			"profiles.url",
			"profiles.external_id",
			"profiles.gender",
			"profiles.birth_date",
			"profiles.person_id",
			"profiles.created_at",
			"profiles.updated_at",
		).
		Joins("JOIN photos ON photos.profile_id = profiles.id").
		Where("photos.index_id IN ?", indexIDs).
		Limit(limit).
		Preload("Photos").
		Find(&profiles).Error

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
