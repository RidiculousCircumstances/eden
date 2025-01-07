package message_processor

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/service/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/infrastructure/queue/message"
	"errors"
)

var (
	ErrPhotoNotFound = errors.New("photo not found by index ID")
)

type TraceFaceMessageProcessor struct {
	faceService  interfaces.FaceService
	photoService interfaces.PhotoService
}

func NewTraceFaceMessageProcessor(faceService interfaces.FaceService, photoService interfaces.PhotoService) consumerIntf.TraceFaceMessageProcessor {
	return &TraceFaceMessageProcessor{
		faceService:  faceService,
		photoService: photoService,
	}
}

func (p *TraceFaceMessageProcessor) Process(ctx context.Context, msg message.TraceFaceMessage) error {
	photoID, err := p.photoService.GetPhotoIdByIndexId(ctx, msg.PhotoId)
	if err != nil {
		return err
	}
	if photoID == 0 {
		return ErrPhotoNotFound
	}

	for _, face := range msg.Faces {
		if err := p.saveFace(ctx, photoID, face); err != nil {
			return err
		}
	}
	return nil
}

func (p *TraceFaceMessageProcessor) saveFace(ctx context.Context, photoId uint, faceItem message.Face) error {
	face := domain.Face{
		PhotoID: photoId,
		Age:     faceItem.Age,
		Sex:     faceItem.Sex,
		Bbox:    faceItem.Bbox,
	}

	return p.faceService.CreateFace(ctx, &face)
}
