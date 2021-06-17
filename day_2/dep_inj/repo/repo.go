package repo

import (
	"database/sql"

	"github.com/Metalscreame/go-training/day_2/dep_inj/entity"
)

type PersonRepository struct {
	database *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func (repository *PersonRepository) FindAll() (res []*entity.Person) {
	// selecting
	return res
}

func (repository *PersonRepository) FindAllBySmth() (res []*entity.Person) {
	// selecting
	return res
}
