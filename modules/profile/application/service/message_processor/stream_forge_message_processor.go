package message_processor

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/service/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/domain/source"
	"eden/modules/profile/infrastructure/queue/message"
	"errors"
	"time"
)

var (
	ErrInvalidSource = errors.New("invalid source alias")
)

type messageProcessor struct {
	profileService interfaces.ProfileService // Сервис для работы с профилями
	photoService   interfaces.PhotoService   // Сервис для работы с фотографиями
}

func NewStreamForgeMessageProcessor(profileService interfaces.ProfileService, photoService interfaces.PhotoService) consumerIntf.StreamForgeMessageProcessor {
	return &messageProcessor{
		profileService: profileService,
		photoService:   photoService,
	}
}

func (p *messageProcessor) Process(ctx context.Context, msg message.StreamForgeMessage) error {
	sourceId, ok := source.GetIDBySourceAlias(msg.SourceAlias)
	if !ok {
		return ErrInvalidSource
	}

	profile := domain.Profile{
		SourceID:   sourceId,
		CityID:     msg.CityID,
		URL:        msg.URL,
		ExternalID: msg.ProfileID,
		Gender:     msg.Gender,
		BirthDate:  msg.BirthDate,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
	if err := p.profileService.CreateOrUpdateProfile(ctx, &profile); err != nil {
		return err
	}

	for _, photoMsg := range msg.Photos {
		photo := domain.Photo{
			ProfileID: profile.ID,
			URL:       photoMsg.PhotoUrl,
			IndexID:   photoMsg.PhotoId,
		}
		if err := p.photoService.CreatePhoto(ctx, &photo); err != nil {
			return err
		}
	}
	return nil
}
