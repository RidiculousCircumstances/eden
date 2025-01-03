package repository

import (
	"context"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
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
	// Используем Find, чтобы избежать ошибки в случае отсутствия записи
	result := r.db.WithContext(ctx).Find(&profile, id)

	if result.Error != nil {
		// Если ошибка не связана с отсутствием записи, возвращаем ошибку
		return nil, result.Error
	}

	// Если профиль не найден (record not found), просто возвращаем nil без ошибки
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &profile, nil
}

func (r *profileRepository) Update(ctx context.Context, profile *domain.Profile) error {
	// Обновляет данные профиля в базе данных
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *profileRepository) Delete(ctx context.Context, id uint) error {
	// Удаляет профиль по его ID
	return r.db.WithContext(ctx).Delete(&domain.Profile{}, id).Error
}
