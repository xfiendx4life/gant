package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/xfiendx4life/gant/pkg/models"
)

type LineStorage interface {
	Create(line *models.Line) error
	UpdateTime(id uuid.UUID, start, finish time.Time) error
	// * mod is a string to choose the way to update staff
	// * del to delete, add to add
	UpdateStaff(id uuid.UUID, mod string, staff ...[]*models.User) error
	Delete(id uuid.UUID) error
}
