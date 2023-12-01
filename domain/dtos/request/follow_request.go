package request

type FollowRequest struct {
	UserIdFollower  uint `json:"user_id_follower"`
	UserIdFollowing uint `json:"user_id_following"`
}
