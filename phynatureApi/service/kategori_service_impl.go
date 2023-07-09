package service

import (
	"phynatureApi/entity"
	"phynatureApi/repository"
)

type kategoriService struct {
	kategoriRepo repository.KategoriRepo
}

func NewKategoriService(kategoriRepo repository.KategoriRepo) KategoriService {
	return &kategoriService{kategoriRepo: kategoriRepo}
}

func (s *kategoriService) CreateKategori(kategori *entity.Kategori) error {
	err := s.kategoriRepo.CreateKategori(kategori)
	if err != nil {
		return err
	}
	return nil
}

func (s *kategoriService) GetAllKategori() ([]entity.Kategori, error) {
	kategori, err := s.kategoriRepo.GetAllKategori()
	if err != nil {
		return nil, err
	}
	return kategori, nil
}
