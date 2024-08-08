package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnect struct {
	DB *gorm.DB
}

func (connect *DbConnect) ConnectDB() error {
	dsn := "host=localhost user=help_user password=help_password123 dbname=helpinghand port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		connect.DB = db
	}
	fmt.Printf("The DB connection : %v\n", db)
	return nil
}

type ConnectionDB interface {
	ConnectDB() error
}
