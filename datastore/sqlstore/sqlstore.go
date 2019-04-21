package sqlstore

import (
	"fmt"

	"github.com/gobuffalo/packr"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
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

func (db *sqlstore) Migrate() {
	migrationSource := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations"),
		Dir: db.conf.Driver,
	}
	migrate.SetTable("migrations")

	migrations, err := migrationSource.FindMigrations()
	if err != nil {
		log.Errorf("Error loading database migrations: %s", err)
	}

	if len(migrations) == 0 {
		log.Fatalf("Missing database migrations")
	}

	n, err := migrate.Exec(db.DB.DB, db.conf.Driver, migrationSource, migrate.Up)
	if err != nil {
		log.Errorf("Error applying database migrations: %s", err)
	}

	if n > 0 {
		log.Infof("Applied %d database migrations!", n)
	}
}
