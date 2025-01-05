package message

type StreamForgeMessage struct {
	SourceAlias string  `json:"source_alias"`
	ProfileID   string  `json:"profile_id"`
	CityID      uint    `json:"city_id"`
	URL         string  `json:"url"`
	Gender      int     `json:"gender"`
	BirthDate   string  `json:"birth_date"`
	Photos      []Photo `json:"photos"`
}

type Photo struct {
	PhotoId  uint32 `json:"photo_id"`
	PhotoUrl string `json:"photo_url"`
}

type TraceFaceMessage struct {
	PhotoId uint32 `json:"photo_id"`
	Faces   []Face `json:"faces_info"`
}

type Face struct {
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
	Bbox string `json:"bbox"`
}
