package service

import (
	"context"
	servIntf "eden/modules/profile/application/usecase/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
)

type photoService struct {
	repo interfaces.PhotoRepository
}

func NewPhotoService(repo interfaces.PhotoRepository) servIntf.PhotoService {
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

func (s *photoService) GetProfileByIndexId(ctx context.Context, indexId uint32) (*domain.Profile, error) {
	return s.repo.GetProfileByPhotoIndexID(ctx, indexId)
}

func (s *photoService) GetProfilesByIndexIds(ctx context.Context, indexIds []uint32) ([]domain.Profile, error) {
	return s.repo.GetProfilesByPhotoIndexIDs(ctx, indexIds)
}
