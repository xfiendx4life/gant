package storage

import (
	"github.com/google/uuid"
	"github.com/xfiendx4life/gant/pkg/models"
)

type UserStorage interface {
	Create(user *models.User) (id string, err error)
	Get(email string) (*models.User, error)
	Delete(email string) error
	Edit(id uuid.UUID, data map[string]string) error
}
