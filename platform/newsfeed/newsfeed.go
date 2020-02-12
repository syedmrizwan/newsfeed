package newsfeed

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type Getter interface {
	GetAll() []Item
}
type Adder interface {
	Add(item Item)
}
type Item struct {
	tableName struct{}  `pg:"item"`
	Title     string    `pg:"title"`
	Post      string    `pg:"post"`
	Stats     StatsType `pg:"stats"`
	CreatedAt time.Time `pg:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `pg:"updated_at" sql:"default:now()"`
}

// CreateItemTable does create Item table to the database
func CreateItemTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Item{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table Item, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table Items created successfully.\n")
	return nil
}

type StatsType struct {
	Views int `json:"views"`
	Likes int `json:"likes"`
}
type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
