package repository

import "phynatureApi/entity"

type KategoriRepo interface {
	CreateKategori(kategori *entity.Kategori) error
	GetAllKategori() ([]entity.Kategori, error)
}
