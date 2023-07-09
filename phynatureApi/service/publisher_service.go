package service

import "phynatureApi/entity"

type PublisherService interface {
	CreatePublisher(publisher *entity.Publisher) error
	GetAllPublisher() ([]entity.Publisher, error)
}
