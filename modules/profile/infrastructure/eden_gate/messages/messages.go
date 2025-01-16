package messages

type ProfileSearchCompletedEvent struct {
	RequestId string    `json:"request_id"`
	Profiles  []Profile `json:"profiles"`
}

type Profile struct {
	ProfileId uint    `json:"profile_id"`
	Url       string  `json:"url"`
	Photos    []Photo `json:"photos"`
}

type Photo struct {
	PhotoId   uint32 `json:"photo_id"`
	ProfileId uint   `json:"profile_id"`
	PhotoUrl  string `json:"photo_url"`
}
