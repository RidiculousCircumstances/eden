package service

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/interfaces"
)

type profileService struct {
	repo interfaces.ProfileRepository
}

func NewProfileService(repo interfaces.ProfileRepository) consumerIntf.ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) CreateOrUpdateProfile(ctx context.Context, profile *domain.Profile) error {
	existingProfile, err := s.repo.GetByID(ctx, profile.ID)
	if err != nil {
		return err
	}
	if existingProfile == nil {
		return s.repo.Create(ctx, profile)
	}
	return s.repo.Update(ctx, profile)
}

func (s *profileService) GetProfileByID(ctx context.Context, id uint) (*domain.Profile, error) {
	return s.repo.GetByID(ctx, id)
}
