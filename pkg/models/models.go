package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Project struct {
	ID     uuid.UUID `json:"id"`
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
	Staff  []*User   `json:"staff"`
}

type Line struct {
	ID       uuid.UUID     `json:"id"`
	Start    time.Time     `json:"start"`
	Duration time.Duration `json:"duration"`
	Project  *Project      `json:"project"`
	Users    []*User       `json:"users"`
	Status   string        `json:"status"`
}
