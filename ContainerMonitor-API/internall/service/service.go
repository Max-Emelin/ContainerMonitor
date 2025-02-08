package service

import (
	"ContainerMonitor-API/internall/model"
	"ContainerMonitor-API/internall/repository"

	"github.com/sirupsen/logrus"
)

type Container interface {
	Create(container model.Container) (int, error)
	CreateOrUpdate(container model.Container) (int, error)
	GetAll() ([]model.Container, error)
	GetById(containerId int) (model.Container, error)
}

type Service struct {
	Container
}

func NewService(repos *repository.Repository) *Service {
	logrus.Debug("NewService - initializing service")
	service := &Service{
		Container: NewContainerService(repos.Container),
	}
	logrus.Debug("NewService - service initialized successfully")

	return service
}
