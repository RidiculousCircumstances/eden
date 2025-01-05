package repository

import (
	"context"
	"eden/modules/profile/domain"
	"gorm.io/gorm"
)

type FaceRepository struct {
	db *gorm.DB
}

func NewFaceRepository(db *gorm.DB) *FaceRepository {
	return &FaceRepository{
		db: db,
	}
}

func (r *FaceRepository) Create(ctx context.Context, face *domain.Face) error {
	return r.db.WithContext(ctx).Create(face).Error
}
