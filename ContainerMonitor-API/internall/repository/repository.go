package repository

import (
	"ContainerMonitor-API/internall/model"

	"github.com/jmoiron/sqlx"
)

type Container interface {
	Create(container model.Container) (int, error)
	GetAll() ([]model.Container, error)
	GetById(containerId int) (model.Container, error)
}

type Repository struct {
	Container
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Container: NewContainerPostgres(db),
	}
}
