package repository

import "phynatureApi/entity"

type PublisherRepo interface {
	CreatePublisher(publisher *entity.Publisher) error
	GetAllPublisher() ([]entity.Publisher, error)
}
