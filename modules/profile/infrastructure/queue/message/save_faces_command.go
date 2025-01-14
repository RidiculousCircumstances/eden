package message

type SaveFacesCommand struct {
	PhotoId uint32 `json:"photo_id"`
	Faces   []Face `json:"faces_info"`
}

type Face struct {
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
	Bbox string `json:"bbox"`
}
