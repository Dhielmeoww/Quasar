package repository

import (
	"phynatureApi/entity"

	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{db: db}
}

func (r *PublisherRepository) CreatePublisher(publisher *entity.Publisher) error {
	err := r.db.Create(publisher).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PublisherRepository) GetAllPublisher() ([]entity.Publisher, error) {
	var publisher []entity.Publisher
	err := r.db.Find(&publisher).Error
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
