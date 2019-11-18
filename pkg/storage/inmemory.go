package storage

import (
	"fmt"
	"sync"
)

func CreateInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{m: make(map[string]string)}
}

type ErrResourceNotFound string

func (e ErrResourceNotFound) Error() string {
	return fmt.Sprintf("Resource with name: %v not found", string(e))
}

type InMemoryStorage struct {
	m   map[string]string
	mux sync.Mutex
}

func (s *InMemoryStorage) Persist(name, content string) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[name] = content;
	return nil
}

func (s *InMemoryStorage) Retrieve(name string) (string, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.m[name] != "" {
		return s.m[name], nil
	} else {
		var err ErrResourceNotFound = ErrResourceNotFound(name)
		return "", err
	}
}

func (s *InMemoryStorage) RetrieveAll() ([]string, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	values := make([]string, len(s.m))
	i := 0
	for _, value := range s.m {
		values[i] = value
		i++
	}
	return values, nil
}

func (s *InMemoryStorage) Delete(name string) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.m[name] != "" {
		delete(s.m, name)
		return nil;
	} else {
		var err ErrResourceNotFound = ErrResourceNotFound(name)
		return err
	}
}

