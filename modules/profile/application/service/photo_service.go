package service

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
)

type photoService struct {
	repo interfaces.PhotoRepository
}

func NewPhotoService(repo interfaces.PhotoRepository) consumerIntf.PhotoService {
	return &photoService{repo: repo}
}

func (s *photoService) CreatePhoto(ctx context.Context, photo *domain.Photo) error {
	return s.repo.Create(ctx, photo)
}

func (s *photoService) GetPhotoIdByIndexId(ctx context.Context, indexId uint32) (uint, error) {
	return s.repo.GetIDByIndexID(ctx, indexId)
}

func (s *photoService) GetPhotosByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error) {
	return s.repo.GetByProfileID(ctx, profileID)
}
