package models

type Account struct {
	Id              int64   `db:"id" json:"id"`
	Name            string  `db:"name" json:"name"`
	CreatedAt       string  `db:"created_at" json:"created_at"`
}
