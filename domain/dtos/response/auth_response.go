package response

type LoginResponse struct {
	AccessToken  string `jaon:"access_token"`
	RefreshToken string `jaon:"refresh_token"`
}
