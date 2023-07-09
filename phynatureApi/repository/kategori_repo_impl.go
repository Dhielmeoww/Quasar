package repository

import (
	"phynatureApi/entity"

	"gorm.io/gorm"
)

type KategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) *KategoriRepository {
	return &KategoriRepository{db: db}
}

func (r *KategoriRepository) CreateKategori(kategori *entity.Kategori) error {
	err := r.db.Create(kategori).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *KategoriRepository) GetAllKategori() ([]entity.Kategori, error) {
	var kategori []entity.Kategori
	err := r.db.Find(&kategori).Error
	if err != nil {
		return nil, err
	}
	return kategori, nil
}
