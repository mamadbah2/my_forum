package models

import (
	"database/sql"
	"errors"
)

type ConnDB struct {
	DB *sql.DB
}

var ErrNoRecord = errors.New("NO RECORDS")
