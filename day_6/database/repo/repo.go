package repo

import (
	"database/sql"
	"fmt"

	"github.com/Metalscreame/go-training/day_6/database/entity"
	_ "github.com/lib/pq"
)

type Repository struct {
	Conn *sql.DB
}

func NewRepository(dsn string) (*Repository, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	r := &Repository{
		Conn: conn,
	}

	if err := r.initTables(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Repository) initTables() error {
	var schema = `
CREATE TABLE IF NOT EXISTS person  (
    first_name text,
    last_name text,
    email text,
    state_code text DEFAULT Null
);

CREATE TABLE IF NOT EXISTS place  (
    country text,
    city text NULL,
    telcode integer
)`
	_, err := r.Conn.Exec(schema)
	return err
}

func (r *Repository) InsertPerson(p *entity.Person) error {
	const query = `INSERT INTO person 
    					(first_name, last_name, email, state_code) 
					VALUES ($1, $2, $3, $4)`

	res, err := r.Conn.Exec(query, p.FirstName, p.LastName, p.Email, p.StateCode)
	if err != nil {
		return err
	}
	fmt.Println(res.RowsAffected())
	return nil
}

func (r *Repository) InsertPlace(p *entity.Place) error {
	const query = `INSERT INTO place (country, telcode) VALUES ($1, $2)`

	res, err := r.Conn.Exec(query, p.Country, p.TelCode)
	if err != nil {
		return err
	}
	fmt.Println(res.RowsAffected())
	return nil
}

func (r *Repository) GetPersons() (res []entity.Person, err error) {
	const query = `SELECT first_name, last_name, email, state_code
						FROM person ORDER BY first_name ASC`

	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := entity.Person{}
		if err := rows.Scan(
			&p.FirstName,
			&p.LastName,
			&p.Email,
			&p.StateCode, //start scan into regular string field and then change to pointer
		); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return
}
