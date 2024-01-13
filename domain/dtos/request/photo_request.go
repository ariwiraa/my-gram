package request

type PhotoRequest struct {
	Caption  string   `json:"caption"`
	PhotoUrl string   `json:"photo_url"`
	Tags     []string `json:"tags"`
}

type UpdatePhotoRequest struct {
	Caption string   `json:"caption"`
	Tags    []string `json:"tags"`
}
