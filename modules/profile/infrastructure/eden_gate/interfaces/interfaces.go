package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/queue/message"
)

type Client interface {
	SendSearchResult(ctx context.Context, msg ProfileSearchCompletedEvent) error
}

type ProfileSearchCompletedEvent struct {
	RequestId string    `json:"request_id"`
	Profiles  []Profile `json:"profiles"`
}

type Profile struct {
	Url    string          `json:"url"`
	Photos []message.Photo `json:"photos"`
}
