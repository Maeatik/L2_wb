package domain

import "time"

// Model стркутура данных, состоящая из айди записи, ее названия и двух дат: начальная и конечная
type Model struct {
	Id       string
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

// NewModel конструктор структуры Model
func NewModel(id, title string, from, to time.Time) *Model {
	return &Model{
		Id:       id,
		Title:    title,
		DateFrom: from,
		DateTo:   to,
	}
}
