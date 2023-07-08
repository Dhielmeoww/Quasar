package entity

type Article struct {
	Id           uint   `gorm:"primaryKey"`
	Title        string `json:"title"`
	Body         string `json:"body_post"`
	Category_id  uint
	Publisher_id uint
}
