package raftd

import (
	"path"
	"errors"
	"encoding/json"
	)

type Store struct {
	Nodes map[string]string  `json:"nodes"`
}

func createStore() *Store{
	s := new(Store)
	s.Nodes = make(map[string]string)
	return s
}

// set the key to value, return the old value if the key exists 
func (s *Store) Set(key string, value string) (string, bool) {

	key = path.Clean(key)

	oldValue, ok := s.Nodes[key]

	if ok {
		s.Nodes[key] = value
		return oldValue, true
	} else {
		s.Nodes[key] = value
		return "", false
	}

}

// get the node of the key
func (s *Store) Get(key string) (string, error) {
	key = path.Clean(key)

	value, ok := s.Nodes[key]

	if ok {
		return value, nil
	} else {
		return "", errors.New("Key does not exist")
	}
}

// delete the key, return the old value if the key exists
func (s *Store) Delete(key string) (string, error) {
	key = path.Clean(key)

	oldValue, ok := s.Nodes[key]

	if ok {
		delete(s.Nodes, key)
		return oldValue, nil
	} else {
		return "", errors.New("Key does not exist")
	}
}

func (s *Store) Save() ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Store) Recovery(state []byte) error {
	err := json.Unmarshal(state, s)
	return err
}