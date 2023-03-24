package posgtresql

import (
	"database/sql"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	*sql.DB
}

var _ database.SQLExecutor = PostgreSQL{}
var _ database.SQLQuerier = PostgreSQL{}
var _ database.SQLPreparation = PostgreSQL{}

func New(conf config.Database) (*PostgreSQL, error) {
	db, err := sql.Open("postgres", conf.DSN)
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{db}, nil
}
