package mappers

import (
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"Spy-Cat-Agency/src/internal/spycats/dtos"

	"github.com/google/uuid"
)

func SpyCatSingleToDto(spycat *models.SpyCat) dtos.SpyCatSingleResponseDto {
	return dtos.SpyCatSingleResponseDto{
		Id:              spycat.Id,
		Name:            spycat.Name,
		ExperienceYears: spycat.ExperienceYears,
		Breed:           spycat.Breed,
		Salary:          spycat.Salary,
		UpdatedAt:       spycat.UpdatedAt,
	}
}

func SpyCatsToDto(spycats []models.SpyCat) []dtos.SpyCatAllResponseDto {

	spycatsResponse := make([]dtos.SpyCatAllResponseDto, 0)

	for _, spycat := range spycats {
		spycatResponse := dtos.SpyCatAllResponseDto{
			Id:              spycat.Id,
			Name:            spycat.Name,
			ExperienceYears: spycat.ExperienceYears,
			Breed:           spycat.Breed,
			Salary:          spycat.Salary,
		}

		spycatsResponse = append(spycatsResponse, spycatResponse)
	}

	return spycatsResponse

}

func CreateDtoToSpyCat(spycatReq dtos.SpyCatCreateRequest) *models.SpyCat {
	return &models.SpyCat{
		Id:              uuid.New(),
		Name:            spycatReq.Name,
		ExperienceYears: spycatReq.ExperienceYears,
		Breed:           spycatReq.Breed,
		Salary:          spycatReq.Salary,
	}
}
