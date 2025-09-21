package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

const MAXSALARY = 9_999_999.99
const MINPOSSIBLESALARY = 0

var ErrNegativeSalary = errors.New("Salary can't be negative")

func (s *spyCatServiceImpl) UpdateSpyCatSalary(ctx context.Context, salaryReq dtos.SalaryRequest, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error) {

	if salaryReq.Salary < MINPOSSIBLESALARY || salaryReq.Salary > MAXSALARY {
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusBadRequest, "Negative salary inputted", ErrNegativeSalary)
	}

	spycat := &models.SpyCat{
		Id:     id,
		Salary: salaryReq.Salary,
	}

	updatedSpyCat, err := s.updater.Update(ctx, spycat)
	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating spy cat's salary in database", err)
	}

	return mappers.SpyCatSingleToDto(updatedSpyCat), nil

}
