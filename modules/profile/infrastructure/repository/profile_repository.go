package repository

import (
	"context"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
	"errors"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

// NewProfileRepository создает новый экземпляр репозитория Profile.
func NewProfileRepository(db *gorm.DB) interfaces.ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) Create(ctx context.Context, profile *domain.Profile) error {
	// Создает новый профиль в базе данных
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *profileRepository) GetByID(ctx context.Context, id uint) (*domain.Profile, error) {
	var profile domain.Profile
	err := r.db.WithContext(ctx).Preload("Photos").First(&profile, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Если профиль не найден, возвращаем nil
	}
	return &profile, err
}

func (r *profileRepository) Update(ctx context.Context, profile *domain.Profile) error {
	// Обновляет данные профиля в базе данных
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *profileRepository) Delete(ctx context.Context, id uint) error {
	// Удаляет профиль по его ID
	return r.db.WithContext(ctx).Delete(&domain.Profile{}, id).Error
}
