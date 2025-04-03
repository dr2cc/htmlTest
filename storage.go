package main

import (
	"errors"
	"fmt"
	"sync"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e Employee) {
	s.Lock()

	e.Id = s.counter
	s.data[e.Id] = *e

	s.counter++
	s.Unlock()

}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, existss := s.data[id]
	if !existss {
		return Employee{}, errors.New("employee with such id doesn`t exist")
	}
	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	//The delete built-in function deletes the element
	//with the specified key (m[key]) from the map.
	//func delete(m map[Type]Type1, key Type)
	delete(s.data, id)
	return nil
}

type dumbStorage struct{}

func NewDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) Insert(e Employee) error {
	fmt.Printf("insert to user with id: %d is done\n", e.Id)
	return nil
}

func (s *dumbStorage) Get(id int) (Employee, error) {
	e := Employee{
		Id: id,
	}

	return e, nil
}

func (s *dumbStorage) Delete(id int) error {
	fmt.Printf("delete to user with id: %d is done\n", id)
	return nil
}
