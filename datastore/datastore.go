package datastore

import "github.com/ysv/pick/datastore/sqlstore"

type Datastore interface {
	InsertPageviews() error
	ListPageviews() error

	Migrate()
	Health() error
}

func NewSQLStore(conf sqlstore.Config) Datastore{
	return sqlstore.New(conf)
}
