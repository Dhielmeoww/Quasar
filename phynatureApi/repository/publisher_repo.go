package repository

import "phynatureApi/entity"

type Publisher_repo interface {
	CreatePublisher(publisher *entity.Publisher) error
	GetAllPublisher() ([]entity.Publisher, error)
}
