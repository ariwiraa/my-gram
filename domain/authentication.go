package domain

type Authentication struct {
	RefreshToken string `gorm:"type:text" json:"refresh_token"`
}
