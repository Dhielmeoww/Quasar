package service

import (
	"phynatureApi/entity"
	"phynatureApi/repository"
)

type publisherService struct {
	publisherRepo repository.PublisherRepository
}

func NewPublisherService(publisherRepo repository.PublisherRepository) PublisherService {
	return &publisherService{publisherRepo: publisherRepo}
}

func (s *publisherService) CreatePublisher(publisher *entity.Publisher) error {
	err := s.publisherRepo.CreatePublisher(publisher)
	if err != nil {
		return err
	}
	return nil
}

func (s *publisherService) GetAllPublisher() ([]entity.Publisher, error) {
	publisher, err := s.publisherRepo.GetAllPublisher()
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
