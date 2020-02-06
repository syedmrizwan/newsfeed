package database

import (
	"github.com/go-pg/pg/v9"
)

var db *pg.DB

func init() {
	//todo add more config to the connection like idle timeout etc
	db = pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "root",
		Password: "root",
		Database: "newsfeeder",
		PoolSize: 10,
	})
}

func GetConnection() *pg.DB {
	return db
}
