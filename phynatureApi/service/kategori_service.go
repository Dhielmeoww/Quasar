package service

import "phynatureApi/entity"

type KategoriService interface {
	CreateKategori(kategori *entity.Kategori) error
	GetAllKategori() ([]entity.Kategori, error)
}
