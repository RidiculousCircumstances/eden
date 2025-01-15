package message

type SearchProfilesCommand struct {
	RequestId string   `json:"request_id"`
	PhotoIds  []uint32 `json:"photo_ids"`
}
