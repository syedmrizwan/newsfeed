package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"newsfeeder/platform/newsfeed"
)

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "root",
		Password: "root",
		Database: "newsfeeder",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	item1 := &newsfeed.Item{
		Title: "news title",
		Post:  "news post",
	}
	err = db.Insert(item1)
	if err != nil {
		panic(err)
	}

	// // Select user by primary key.
	// user := &newsfeed.Item{Id: user1.Id}
	// err = db.Select(user)
	// if err != nil {
	//     panic(err)
	// }

	// Select all users.
	var items []newsfeed.Item
	err = db.Model(&items).Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(items)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*newsfeed.Item)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ExampleDB_Model()
}
