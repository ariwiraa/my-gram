package response

import "time"

type PhotoResponse struct {
	Id            string     `json:"id"`
	Caption       string     `json:"caption"`
	PhotoUrl      string     `json:"photo_url"`
	PhotoTags     []string   `json:"photo_tags,omitempty"`
	TotalLikes    int64      `json:"total_likes"`
	TotalComments int64      `json:"total_comments"`
	Username      string     `json:"username"`
	CreatedAt     *time.Time `json:"created_at"`
}
