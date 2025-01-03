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

func (s *photoService) CreateOrUpdatePhoto(ctx context.Context, photo *domain.Photo) error {
	existingPhoto, err := s.repo.GetByID(ctx, photo.ID)
	if err != nil {
		return err
	}
	if existingPhoto == nil {
		return s.repo.Create(ctx, photo)
	}
	return s.repo.Update(ctx, photo)
}

func (s *photoService) GetPhotosByProfileID(ctx context.Context, profileID uint) ([]domain.Photo, error) {
	return s.repo.GetByProfileID(ctx, profileID)
}
