package models

type Animal struct {
	ID    int    `json:"id" db:"id" gorm:"primary_key"`
	Name  string `json:"name" db:"name"`
	Class string `json:"class" db:"class"`
	Legs  string `json:"legs" db:"legs"`
}
