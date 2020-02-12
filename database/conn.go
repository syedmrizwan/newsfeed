package database

import (
	"newsfeeder/env"
	"newsfeeder/platform/newsfeed"

	"github.com/go-pg/pg/v9"
)

var db *pg.DB

func init() {
	//todo add more config to the connection like idle timeout etc
	db = pg.Connect(&pg.Options{
		Addr:     env.Env.GetAddr(),
		User:     env.Env.DbUsername,
		Password: env.Env.DbPassword,
		Database: env.Env.DbName,
		PoolSize: env.Env.DbPoolSize,
	})
	newsfeed.CreateItemTable(db)
}

// GetConnection returns pg connection
func GetConnection() *pg.DB {
	return db
}
