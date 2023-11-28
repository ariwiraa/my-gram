package domain

type Tag struct {
	ID    uint    `gorm:"primarykey" json:"id"`
	Name  string  `json:"name"`
	Photo []Photo `gorm:"many2many:photo_tags" json:"photo,omitempty"`
}
