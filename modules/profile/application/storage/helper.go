package storage

type Helper struct {
	clientBucket string
}

func NewHelper(clientBucket string) *Helper {
	return &Helper{clientBucket: clientBucket}
}

func (h *Helper) GetObjectName(fileId string) string {
	return h.clientBucket + "/" + fileId
}
