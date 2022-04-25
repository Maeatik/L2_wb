package domain

import (
	"sync"
	"time"
)

//Интерфейс для обработки запросов
type StorageInterface interface {
	GetAll() ([]*Model, error)
	GetById(id string) (*Model, error)
	GetByPeriod(from, to time.Time) ([]*Model, error)
	Save(event *Model) error
	Delete(event *Model) error
}

// Storage тип хранения, который может лочить каналы и имеет в себе сведения о теле запроса
type Storage struct {
	sync.Mutex
	models map[string]*Model
}

func NewStorage() *Storage {
	eventStorage := make(map[string]*Model)
	return &Storage{models: eventStorage}
}

//так как мапы не потоко-безопасные - везде используем мьютекс

// GetAll Получить все добавленные объекты
func (s *Storage) GetAll() ([]*Model, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	res := make([]*Model, 0, len(s.models))

	for _, event := range s.models {
		res = append(res, event)
	}
	return res, nil
}

// GetById получить добавленный объект по айди
func (s *Storage) GetById(id string) (*Model, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if model, ok := s.models[id]; ok {
		return model, nil
	}
	return nil, nil
}

// GetByPeriod получить все добавленные объекты по айди по дате
func (s *Storage) GetByPeriod(from, to time.Time) ([]*Model, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	res := make([]*Model, 0)

	for _, model := range s.models {
		if model.DateFrom.Before(to) && model.DateTo.After(from) {
			res = append(res, model)
		}
	}
	return res, nil
}

// Save добавить в хранилище новую модель
func (s *Storage) Save(model *Model) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.models[model.Id] = model

	return nil
}

// Delete удалить из хранилища по айди модель
func (s *Storage) Delete(model *Model) error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	delete(s.models, model.Id)

	return nil
}
