package message

type SaveProfileCommand struct {
	SourceAlias string  `json:"source_alias"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
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
