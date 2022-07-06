package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDb struct {
	DB *gorm.DB
}

//SetupDb sets up database and auto migrate schema
func (pdb *PostgresDb) SetupDb(host, user, password, dbName, port string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pdb.DB = db

	dberr := pdb.DB.AutoMigrate()
	if dberr != nil {
		log.Fatal(dberr)
	}
	return nil
}
