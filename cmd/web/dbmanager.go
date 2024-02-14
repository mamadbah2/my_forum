package main

import (
	"database/sql"
	"os"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) createTable() error {
	fileByte, err := os.ReadFile("./db/migration.sql")
	if err != nil {
		return err
	}
	sqlStatements := string(fileByte)
	_, err = app.connDB.DB.Exec(sqlStatements)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) setData() error {
	content, err := os.ReadFile("./db/data.sql")
	if err != nil {
		return err
	}
	statements := string(content)
	_, err = app.connDB.DB.Exec(statements)
	if err != nil {
		return err
	}
	return nil
}
