package usecase

type AuthenticationUsecase interface {
	Add(token string) error
	ExistsByRefreshToken(token string) error
	Delete(token string) error
}
