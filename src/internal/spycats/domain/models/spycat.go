package models

import (
	"time"

	"github.com/google/uuid"
)

type SpyCat struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float64   `json:"salary"`
	CreatedAt       time.Time
	UpdatedAt       time.Time `json:"updated_at"`
}
