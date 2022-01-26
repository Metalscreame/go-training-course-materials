package main

import (
	"fmt"
	"os"

	"github.com/Metalscreame/go-training/day_6/database/entity"
	_ "github.com/jackc/pgx/v4"

	"github.com/Metalscreame/go-training/day_6/database/repo"
)
// SQL_DSN=postgres://postgres:secret@localhost:5432/test?sslmode=disable;SQL_DRIVER=pgx

func main() {
	dsn := os.Getenv("SQL_DSN")
	repo, err := repo.NewRepository(dsn)
	if err != nil {
		fmt.Println(err)
	}

	tx, _ := repo.Conn.Begin()
	/*
		tx0, err := repo.Conn.BeginTx(context.Background(), &sql.TxOptions{
			Isolation: sql.LevelReadUncommitted,
		})
	*/

	err = repo.InsertPerson(&entity.Person{
		FirstName: "Roman",
		LastName:  "Kosyi",
		Email:     "me@you.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return
	}



	persons, err := repo.GetPersons()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(persons)
}
