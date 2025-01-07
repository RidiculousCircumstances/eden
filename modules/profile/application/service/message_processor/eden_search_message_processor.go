package message_processor

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/service/interfaces"
	"eden/modules/profile/domain"
	edenGateIntf "eden/modules/profile/infrastructure/eden_gate/interfaces"
	"eden/modules/profile/infrastructure/queue/message"
)

type EdenSearchMessageProcessor struct {
	photoService interfaces.PhotoService
	publisher    consumerIntf.EdenGateSearchResultPublisher
}

func NewEdenSearchMessageProcessor(
	photoService interfaces.PhotoService,
	publisher consumerIntf.EdenGateSearchResultPublisher,
) consumerIntf.EdenSearchMessageProcessor {
	return &EdenSearchMessageProcessor{
		photoService: photoService,
		publisher:    publisher,
	}
}

func (p *EdenSearchMessageProcessor) Process(ctx context.Context, msg message.SearchProfileMessage) error {
	profiles, err := p.photoService.GetProfilesByIndexIds(ctx, msg.PhotoIds)
	if err != nil {
		return err
	}

	searchResultMsg := buildSearchResultMessage(msg.RequestId, profiles)

	err = p.publisher.Publish(ctx, searchResultMsg)
	if err != nil {
		return err
	}

	return nil
}

func buildSearchResultMessage(requestId string, profiles []domain.Profile) edenGateIntf.EdenGateSearchResultMessage {
	return edenGateIntf.EdenGateSearchResultMessage{
		RequestId: requestId,
		Profiles:  buildProfiles(profiles),
	}
}

func buildProfiles(profiles []domain.Profile) []edenGateIntf.Profile {
	result := make([]edenGateIntf.Profile, len(profiles))
	for i, profile := range profiles {
		result[i] = edenGateIntf.Profile{
			Url:    profile.URL,
			Photos: buildPhotos(profile.Photos),
		}
	}
	return result
}

func buildPhotos(photos []*domain.Photo) []message.Photo {
	result := make([]message.Photo, len(photos))
	for i, photo := range photos {
		result[i] = message.Photo{
			PhotoId:  photo.IndexID,
			PhotoUrl: photo.URL,
		}
	}
	return result
}
