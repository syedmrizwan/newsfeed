package newsfeed

type Getter interface {
	GetAll() []Item
}
type Adder interface {
	Add(item Item)
}
type Item struct {
	tableName struct{}  `pg:"item"`
	Title     string    `json:"title", pg:"title"`
	Post      string    `json:"post", pg:"post"`
	Stats     StatsType `pg:"stats"`
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
