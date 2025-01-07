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
