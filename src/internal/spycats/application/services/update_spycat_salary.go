package services

import (
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

const MAXSALARY = 9_999_999.99

func (s *spyCatServiceImpl) UpdateSpyCatSalary(ctx context.Context, salaryReq dtos.SalaryRequest, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error) {

	if salaryReq.Salary < 0 || salaryReq.Salary > MAXSALARY {
		return dtos.SpyCatSingleResponseDto{}, errors.New("Salary can't be negative")
	}

	spycat := &models.SpyCat{
		Id:     id,
		Salary: salaryReq.Salary,
	}

	updatedSpyCat, err := s.updater.Update(ctx, spycat)
	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, err
	}

	return mappers.SpyCatSingleToDto(updatedSpyCat), nil

}
