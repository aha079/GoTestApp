package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    Connection *sql.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    return &Database{Connection: db}, nil
}

func (d *Database) Close() {
    d.Connection.Close()
}
