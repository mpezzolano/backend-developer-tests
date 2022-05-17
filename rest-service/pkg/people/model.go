package people

import (
	"github.com/gofrs/uuid"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/storage"
)

type PersonDTO struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	PhoneNumber string
}

func fromStorage(person *storage.Person) PersonDTO {
	return PersonDTO(*person)
}

//service
type Person struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
}

type PersonFilter struct {
	FirstName   string
	LastName    string
	PhoneNumber string
}

func (p PersonFilter) IsEmpty() bool {
	return p.FirstName == "" && p.LastName == "" && p.PhoneNumber == ""
}

func fromDto(dto PersonDTO) Person {
	return Person(dto)
}

func fromDtoArr(dtoArr []PersonDTO) []Person {
	result := make([]Person, len(dtoArr))

	for i, dto := range dtoArr {
		result[i] = fromDto(dto)
	}

	return result
}

// transport
type GetPeopleResponse struct {
	Data  []Person `json:"data"`
	Total int      `json:"total"`
}

type GetPeopleRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone_number"`
}

type GetPersonByIDRequest struct {
	ID uuid.UUID
}
