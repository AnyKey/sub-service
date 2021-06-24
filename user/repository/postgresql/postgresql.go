package postgresql

import (
	"database/sql"
	"github.com/AnyKey/sub-service/user"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	db *sql.DB
}

// New will create new an Repository object representation of user.Repository interface
func New(db *sql.DB) user.Repository {
	return &Repository{db}
}

// GetOption get user options for api request
func (nr *Repository) GetOption() error {
	// TODO: implementation
	log.Infoln("User option received")

	return nil
}
