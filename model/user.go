package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type User struct {
	tableName struct{}  `pg:"user"`
	UserName  string    `pg:"username" sql:",pk"`
	Password  string    `pg:"password"`
	CreatedAt time.Time `pg:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `pg:"updated_at" sql:"default:now()"`
}

// CreateItemTable does create Item table to the database
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&User{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table User, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table User created successfully.\n")
	return nil
}
