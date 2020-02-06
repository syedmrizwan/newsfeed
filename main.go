package main

import (
	"fmt"
	"newsfeeder/database"
	"newsfeeder/platform/newsfeed"
)

func ExampleDB_Model() {
	db := database.GetConnection()
	defer db.Close()

	item1 := &newsfeed.Item{
		Title: "news2 title",
		Post:  "news 2post",
	}
	err := db.Insert(item1)
	if err != nil {
		panic(err)
	}

	// Select all items.
	var items []newsfeed.Item
	err = db.Model(&items).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(items)
}

func main() {
	ExampleDB_Model()
}
