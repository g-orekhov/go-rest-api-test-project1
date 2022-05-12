package db

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var errDataBase error = errors.New("data base common error")
var errNotFound error = errors.New("records by request was not found")

type DB struct {
	pool *gorm.DB
}

var database *DB

func (db *DB) ConnectDB() error {
	if database == nil {
		database = db
	}
	//TODO: make configurations
	dialector := postgres.Open(os.Getenv("DB_URL"))
	var err error
	//TODO: можно ли так делать? Какая часть должна быть общая, а какая одтдельная для каждого запроса?
	db.pool, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func GetDB() *DB {
	return database
}
