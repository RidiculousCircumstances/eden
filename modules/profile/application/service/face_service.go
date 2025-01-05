package service

import (
	"context"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
)

type FaceService struct {
	faceRepo interfaces.FaceRepository
}

func NewFaceService(faceRepo interfaces.FaceRepository) *FaceService {
	return &FaceService{faceRepo}
}

func (s *FaceService) CreateFace(ctx context.Context, face *domain.Face) error {
	return s.faceRepo.Create(ctx, face)
}
