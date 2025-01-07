package interfaces

import (
	"context"
	"eden/modules/profile/infrastructure/queue/message"
)

type Client interface {
	SendSearchResult(ctx context.Context, msg EdenGateSearchResultMessage) error
}

type EdenGateSearchResultMessage struct {
	RequestId string    `json:"request_id"`
	Profiles  []Profile `json:"profiles"`
}

type Profile struct {
	Url    string          `json:"url"`
	Photos []message.Photo `json:"photos"`
}
