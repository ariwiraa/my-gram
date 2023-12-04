package response

import "github.com/ariwiraa/my-gram/domain"

type UserProfileResponse struct {
	Username   string         `json:"username"`
	PostsCount int64          `json:"posts_count"`
	Follower   int64          `json:"follower"`
	Following  int64          `json:"following"`
	Posts      []domain.Photo `json:"posts"`
}
