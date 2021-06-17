package service

import (
	"github.com/Metalscreame/go-training/day_2/dep_inj/entity"
)


// general and consumer interfaces
// duck typing
// helps with creating mocks
type Repository interface {
	FindAll() (res []*entity.Person)
}

type PersonService struct {
	repo  Repository
}

func NewPersonService(r Repository) *PersonService {
	return &PersonService{repo: r}
}
