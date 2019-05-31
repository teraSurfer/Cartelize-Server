package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
)

// DB type
type DB struct {
	*sql.DB
}

const (
	// DbUser name
	DbUser = "postgres"

	// DbPassword password
	DbPassword = "admin"

	// DbName database name
	DbName = "cartelizedb"
)

var db *gorp.DbMap

// Init database
func Init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser,
		DbPassword,
		DbName,
	)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB provides a database map on successful connection.
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}
	return dbmap, nil
}

// GetDB returns DB map on success.
func GetDB() *gorp.DbMap {
	return db
}
