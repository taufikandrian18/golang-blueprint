package service

import (
	"database/sql"

	"gitlab.com/wit-id/project-latihan/toolkit/config"
)

type EmployeeService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

func NewEmployeeService(
	mainDB *sql.DB,
	cfg config.KVStore,
) *EmployeeService {
	return &EmployeeService{
		mainDB: mainDB,
		cfg:    cfg,
	}
}
