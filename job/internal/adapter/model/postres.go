package model

import "database/sql"

type DBhandler interface {
	Beginx() (*sql.Tx, error)
}
