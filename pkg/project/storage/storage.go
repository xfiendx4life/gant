package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/xfiendx4life/gant/pkg/models"
)

type ProjectStorage interface {
	Create(project *models.Project) error
	Get(id uuid.UUID) (*models.Project, error)
	UpdateDates(id uuid.UUID, start, finish time.Time) error
	// * mod is a string to choose the way to update staff
	// * del to delete, add to add
	UpdateStaff(id uuid.UUID, mod string, staffToChange ...[]models.User)
	// TODO: think about transaction to delete all the lines connected to the project
	Delete(id uuid.UUID) error
}