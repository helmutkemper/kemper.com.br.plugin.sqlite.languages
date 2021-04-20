package main

import (
	"database/sql"
)

type SQLite3 struct {
	Database *sql.DB
}
