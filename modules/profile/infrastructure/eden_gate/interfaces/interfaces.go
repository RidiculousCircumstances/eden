package interfaces

import (
	"eden/modules/profile/infrastructure/queue/message"
)

type ProfileSearchCompletedEvent struct {
	RequestId string    `json:"request_id"`
	Profiles  []Profile `json:"profiles"`
}

type Profile struct {
	Url    string          `json:"url"`
	Photos []message.Photo `json:"photos"`
}
