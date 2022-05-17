package people

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/common"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/storage"
)

type Repository interface {
	GetAll(ctx context.Context) ([]PersonDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (PersonDTO, error)
	Find(ctx context.Context, firstName, lastName, phone string) ([]PersonDTO, error)
}

type repository struct {
	st storage.Storage
}

func NewRepository(st storage.Storage) Repository {
	return &repository{st: st}
}

func (r repository) GetAll(_ context.Context) ([]PersonDTO, error) {
	ppl := r.st.AllPeople()
	result := make([]PersonDTO, len(ppl))
	for i, person := range ppl {
		result[i] = fromStorage(person)
	}

	return result, nil
}

func (r repository) GetByID(_ context.Context, id uuid.UUID) (PersonDTO, error) {
	person, err := r.st.FindPersonByID(id)
	if err != nil {
		return PersonDTO{}, errors.New(common.ErrStorage.Error())
	}

	return fromStorage(person), nil
}

func (r repository) Find(_ context.Context, firstName, lastName, phone string) ([]PersonDTO, error) {
	var searchArr []*storage.Person

	if firstName != "" || lastName != "" {
		searchArr = r.st.FindPeopleByName(firstName, lastName)
	}

	if phone != "" {
		searchArr = append(searchArr, r.st.FindPeopleByPhoneNumber(phone)...)
	}

	var result []PersonDTO
	resultMap := make(map[uuid.UUID]bool)
	for _, person := range searchArr {
		if resultMap[person.ID] {
			continue
		}

		resultMap[person.ID] = true
		result = append(result, fromStorage(person))
	}

	return result, nil
}
