package dao

type Website struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Domain string `db:"domain"`
}
