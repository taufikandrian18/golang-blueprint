package service

import (
	"database/sql"

	"gitlab.com/wit-id/project-latihan/toolkit/config"
)

// PostCategoryService ...
type AuthenticationService struct {
	mainDB *sql.DB
	cfg    config.KVStore
}

// NewPostCategoryService ...
func NewAuthenticationService(mainDb *sql.DB, cfg config.KVStore) *AuthenticationService {
	return &AuthenticationService{
		mainDB: mainDb,
		cfg:    cfg,
	}
}
