package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func InitDB(psqlConfig string) (*Postgres, error) {
	var err error

	tempDB, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil {
		fmt.Print("connection error: ", err.Error())
		return nil, err
	}

	return &Postgres{
		db: tempDB,
	}, nil
}
