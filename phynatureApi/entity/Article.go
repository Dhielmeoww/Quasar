package entity

type Article struct {
	Id          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Body        string `json:"body_post"`
	KategoriID  uint
	Kategori    Kategori `gorm:"foreignKey:KategoriID"`
	PublisherID uint
	Publisher   Publisher `gorm:"foreignKey:PublisherID"`
}
