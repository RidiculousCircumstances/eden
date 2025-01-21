package usecase

import (
	"context"
	consumerIntf "eden/modules/profile/application/consumer/interfaces"
	"eden/modules/profile/application/usecase/interfaces"
	"eden/modules/profile/domain"
	"eden/modules/profile/infrastructure/queue/message"
	"errors"
)

var (
	ErrPhotoNotFound = errors.New("photo not found by index ID")
)

type SaveFaceInfo struct {
	faceService  interfaces.FaceService
	photoService interfaces.PhotoService
}

func NewSaveFaceInfo(faceService interfaces.FaceService, photoService interfaces.PhotoService) consumerIntf.SaveFaceInfo {
	return &SaveFaceInfo{
		faceService:  faceService,
		photoService: photoService,
	}
}

func (p *SaveFaceInfo) Process(ctx context.Context, msg message.SaveFacesCommand) error {
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

func (p *SaveFaceInfo) saveFace(ctx context.Context, photoId uint, faceItem message.Face) error {
	face := domain.Face{
		PhotoID: photoId,
		Age:     faceItem.Age,
		Sex:     faceItem.Sex,
		Bbox:    faceItem.Bbox,
	}

	return p.faceService.CreateFace(ctx, &face)
}
