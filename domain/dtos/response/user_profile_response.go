package response

import "github.com/ariwiraa/my-gram/domain"

type UserProfileResponse struct {
	Username   string         `json:"username"`
	PostsCount uint           `json:"posts_count"`
	Follower   uint           `json:"follower"`
	Following  uint           `json:"following"`
	Posts      []domain.Photo `json:"posts"`
}
