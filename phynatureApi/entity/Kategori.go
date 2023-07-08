package entity

type Kategori struct {
	Id            uint   `gorm:"primaryKey"`
	Category_name string `json:"title"`
}
