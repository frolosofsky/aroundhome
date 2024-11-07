package pg

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Driver struct {
	db *sql.DB
}

func NewDriver(conn string) (*Driver, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &Driver{db}, nil
}

func (d *Driver) Health() error {
	var res int
	return d.db.QueryRow("select 1").Scan(&res)
}
