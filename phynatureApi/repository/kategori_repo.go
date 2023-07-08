package repository

import "phynatureApi/entity"

type Kategori_repo interface {
	CreateKategori(kategori *entity.Kategori) error
	GetAllKategori() ([]entity.Kategori, error)
}
