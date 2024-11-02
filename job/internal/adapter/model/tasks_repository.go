package model

import "database/sql"

type DBhandler interface {
	Beginx() (*sql.Tx, error)
}

type Task struct {
	ID   int
	Name string
}
