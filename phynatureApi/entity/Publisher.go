package entity

type Publisher struct {
	Id             uint   `gorm:"primaryKey"`
	Publisher_name string `json:"Publisher"`
}
