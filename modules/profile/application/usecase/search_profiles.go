package usecase

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/usecase/interfaces"
	"eden/modules/profile/domain"
	edenMsg "eden/modules/profile/infrastructure/eden_gate/messages"
	"eden/modules/profile/infrastructure/queue/message"
	loggerIntf "eden/shared/logger/interfaces"
	"go.uber.org/zap"
)

type SearchProfiles struct {
	photoService interfaces.PhotoService
	publisher    consumerIntf.EdenGateSearchResultPublisher
	logger       loggerIntf.Logger
}

func NewSearchProfiles(
	photoService interfaces.PhotoService,
	publisher consumerIntf.EdenGateSearchResultPublisher,
	logger loggerIntf.Logger,
) consumerIntf.SearchProfiles {
	return &SearchProfiles{
		photoService: photoService,
		publisher:    publisher,
		logger:       logger,
	}
}

func (p *SearchProfiles) Process(ctx context.Context, msg message.SearchProfilesCommand) error {
	profiles, err := p.photoService.GetProfilesByIndexIds(ctx, msg.PhotoIds)
	if err != nil {
		return err
	}

	searchResultMsg := buildSearchResultMessage(msg.RequestId, profiles)

	err = p.publisher.Publish(ctx, searchResultMsg)
	if err != nil {
		return err
	}

	p.logger.Info("[SearchProfiles] message processed successfully: ", zap.String("requestId", msg.RequestId))
	return nil
}

func buildSearchResultMessage(requestId string, profiles []domain.Profile) edenMsg.ProfileSearchCompletedEvent {
	return edenMsg.ProfileSearchCompletedEvent{
		RequestId: requestId,
		Profiles:  buildProfiles(profiles),
	}
}

func buildProfiles(profiles []domain.Profile) []edenMsg.Profile {
	result := make([]edenMsg.Profile, len(profiles))
	for i, profile := range profiles {
		result[i] = edenMsg.Profile{
			ProfileId: profile.ID,
			Url:       profile.URL,
			Photos:    buildPhotos(profile.Photos),
		}
	}
	return result
}

func buildPhotos(photos []*domain.Photo) []edenMsg.Photo {
	result := make([]edenMsg.Photo, len(photos))
	for i, photo := range photos {
		result[i] = edenMsg.Photo{
			PhotoId:   photo.IndexID,
			ProfileId: photo.ProfileID,
			PhotoUrl:  photo.URL,
		}
	}
	return result
}
