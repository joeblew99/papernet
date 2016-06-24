package database

import (
	"github.com/bobinette/papernet/models"
)

type DB interface {
	Get(int) (*models.Paper, error)
	List() ([]*models.Paper, error)

	Insert(*models.Paper) error
	Update(*models.Paper) error

	Delete(int) error

	Close() error
}
