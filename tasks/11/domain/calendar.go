package domain

import "time"

// CalendarInterface Полная аналогия с хранилищем
type CalendarInterface interface {
	GetEvents() ([]*Model, error)
	GetEventByID(id string) (*Model, error)
	AddEvent(id, title string, from, to time.Time) (*Model, error)
	UpdateEvent(event *Model) error
	RemoveEvent(event *Model) error
}

type Calendar struct {
	Title   string
	storage StorageInterface
}

func NewCalendar(title string, s StorageInterface) *Calendar {
	return &Calendar{
		Title:   title,
		storage: s,
	}
}

func (c *Calendar) GetEvents() ([]*Model, error) {
	return c.storage.GetAll()
}

func (c *Calendar) GetEventById(id string) (*Model, error) {
	return c.storage.GetById(id)
}

func (c *Calendar) AddEvent(id, title string, from, to time.Time) (*Model, error) {
	model := NewModel(id, title, from, to)
	if err := c.storage.Save(model); err != nil {
		return nil, err
	}
	return model, nil
}

func (c *Calendar) UpdateEvent(model *Model) error {
	return c.storage.Save(model)
}

func (c *Calendar) RemoveEvent(model *Model) error {
	return c.storage.Delete(model)
}
