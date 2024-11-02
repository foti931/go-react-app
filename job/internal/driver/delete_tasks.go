package driver

import (
	"database/sql"
	"go-job/pkg/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func DeleteTasks() error {
	// delete tasks
	cfg := config.Get()

	// -postgres
	db, err := sql.OpenDB(pgdriver., cfg.Postgres.Dsn)
	if err != nil {
		return err
	}
}
