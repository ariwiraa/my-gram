package domain

type PhotoTags struct {
	PhotoId string `json:"photo_id"`
	TagId   uint   `json:"tag_id"`
}

func (PhotoTags) TableName() string {
	return "photo_tags"
}
