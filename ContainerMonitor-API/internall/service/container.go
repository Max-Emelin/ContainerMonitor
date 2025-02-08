package service

import (
	"ContainerMonitor-API/internall/model"
	"ContainerMonitor-API/internall/repository"

	"github.com/sirupsen/logrus"
)

type ContainerService struct {
	repo repository.Container
}

func NewContainerService(repo repository.Container) *ContainerService {
	logrus.Debug("Creating a new instance of ContainerService")
	return &ContainerService{
		repo: repo,
	}
}

func (s *ContainerService) Create(container model.Container) (int, error) {
	id, err := s.repo.Create(container)
	if err != nil {
		logrus.Errorf("Error creating container: %v", err)
		return 0, err
	}

	return id, nil
}

func (s *ContainerService) GetAll() ([]model.Container, error) {
	containers, err := s.repo.GetAll()
	if err != nil {
		logrus.Errorf("Error fetching containers: %v", err)
		return nil, err
	}

	return containers, nil
}

func (s *ContainerService) GetById(containerId int) (model.Container, error) {
	container, err := s.repo.GetById(containerId)
	if err != nil {
		logrus.Errorf("Error fetching container with Id %d: %v", containerId, err)
		return model.Container{}, err
	}

	return container, nil
}
