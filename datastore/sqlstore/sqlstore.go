package sqlstore

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqlstore struct {
	*sqlx.DB
	conf *Config
}

func New(conf Config) *sqlstore{
	conn, err := sqlx.Connect(conf.Driver, conf.ForConnect())
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	}
	ds := sqlstore{conn, &conf}

	return &ds
}

func (*sqlstore) InsertPageviews() error{
	return nil
}

func (*sqlstore) ListPageviews() error{
	return nil
}

func (*sqlstore) Health() error{
	return nil
}
