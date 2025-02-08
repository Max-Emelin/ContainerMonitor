package repository

import (
	"ContainerMonitor-API/internall/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ContainerPostgres struct {
	db *sqlx.DB
}

func NewContainerPostgres(db *sqlx.DB) *ContainerPostgres {
	return &ContainerPostgres{db: db}
}

func (r *ContainerPostgres) Create(container model.Container) (int, error) {
	logrus.Debugf("CreateContainer - IPAddress: %s, PingTime: %s, LastChecked: %s.", container.IPAddress, container.PingTime, container.LastChecked)

	var containerId int
	createContainerQuery := fmt.Sprintf(`INSERT INTO %s (ip_address, ping_time, last_checked)  
										VALUES ($1, $2, $3) 
										RETURNING id`, containersTable)
	err := r.db.QueryRow(createContainerQuery, container.IPAddress, container.PingTime, container.LastChecked).Scan(&containerId)
	if err != nil {
		logrus.Errorf("Error while creating container: %s", err)
		return 0, err
	}

	return containerId, nil
}

func (r *ContainerPostgres) GetAll() ([]model.Container, error) {
	logrus.Debugf("Getting all containers.")

	var containers []model.Container
	getAllContainersQuery := fmt.Sprintf(`SELECT id, ip_address, ping_time, last_checked
										FROM %s`,
		containersTable)
	if err := r.db.Select(&containers, getAllContainersQuery); err != nil {
		logrus.Errorf("Error while getting all containers: %s", err)
		return nil, err
	}

	return containers, nil
}

func (r *ContainerPostgres) GetById(containerId int) (model.Container, error) {
	logrus.Debugf("Getting container with Id = %d.", containerId)

	var container model.Container
	getContainerQuery := fmt.Sprintf(`SELECT id, ip_address, ping_time, last_checked
						FROM %s 
						WHERE id = $1`,
		containersTable)
	if err := r.db.Get(&container, getContainerQuery, containerId); err != nil {
		logrus.Errorf("Error while getting container with Id = %d: %s", containerId, err)
		return container, err
	}

	return container, nil
}

func (r *ContainerPostgres) CreateOrUpdate(container model.Container) (int, error) {
	logrus.Debugf("Create or update container - IPAddress: %s, PingTime: %s, LastChecked: %s.", container.IPAddress, container.PingTime, container.LastChecked)

	var containerId int
	createOrUpdateQuery := fmt.Sprintf(`INSERT INTO %s (ip_address, ping_time, last_checked)  
        							VALUES ($1, $2, $3)
        							ON CONFLICT (ip_address) DO UPDATE 
        							SET ping_time = EXCLUDED.ping_time, last_checked = EXCLUDED.last_checked 
        							RETURNING id`,
		containersTable)
	err := r.db.QueryRow(createOrUpdateQuery, container.IPAddress, container.PingTime, container.LastChecked).Scan(&containerId)
	if err != nil {
		logrus.Errorf("Error while creating/updating container: %s", err)
		return 0, err
	}

	return containerId, nil
}
