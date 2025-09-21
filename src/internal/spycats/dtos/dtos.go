package dtos

import (
	"time"

	"github.com/google/uuid"
)

type SpyCatSingleResponseDto struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float64   `json:"salary"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SpyCatAllResponseDto struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float64   `json:"salary"`
}

type SpyCatCreateRequest struct {
	Name            string  `json:"name" binding:"required"`
	ExperienceYears int     `json:"experience_years" binding:"required"`
	Breed           string  `json:"breed" binding:"required"`
	Salary          float64 `json:"salary" binding:"required"`
}

type BreedName struct {
	Name string `json:"name"`
}

type SalaryRequest struct {
	Salary float64 `json:"salary" binding:"required"`
}

type DeletedIds struct {
	Ids []uuid.UUID `json:"ids" binding:"required"`
}
