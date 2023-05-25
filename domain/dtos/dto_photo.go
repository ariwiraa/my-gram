package dtos

type PhotoRequest struct {
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url"`
}
