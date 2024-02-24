package kvstore

import (
	"fmt"
	"os"
)

type KeyValue struct {
	Key      string
	Value    string
	Revision int
}

type Store struct {
	File string
	Data []KeyValue
}

func OpenDefault() (Store, error) {
	var sf = "global.db"
	return Open(sf)
}

func Open(file string) (Store, error) {
	pairs := make([]KeyValue, 0)
	s := Store {
		File: file,
		Data: pairs,
	}
	// TODO: Unless it's permission denied accessing/creating
	// is it safe to ignore the error?
	s, err := s.readdb()
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s Store) readdb() (Store, error) {
	f, err := os.OpenFile(s.File, os.O_RDONLY|os.O_CREATE, 0700)
	if err != nil {
		return Store{}, err
	}
	defer f.Close()

	for {
		var key, value string
		var revision int
		_, err := fmt.Fscanf(f, "%v\n%v\n%v\n", &key, &value, &revision)
		if err != nil {
			break
		}

		pair := KeyValue{
			Key:      key,
			Value:    value,
			Revision: revision,
		}
		s.Data = append(s.Data, pair)
	}
	return s, nil
}

func (s Store) writedb() error {
	f, err := os.OpenFile(s.File, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf string
	for _, pair := range s.Data {
		buf += fmt.Sprintf("%v\n%v\n%v\n", pair.Key, pair.Value, pair.Revision)
	}
	f.Write([]byte(buf))
	return nil
}

func (s Store) Retrieve(key string) string {
	i, err := s.FindKV(key)
	if err != nil {
		return ""
	}
	return s.Data[i].Value
}

func (s Store) FindKV(key string) (int, error) {
	for i, r := range s.Data {
		if r.Key == key {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

func (s Store) Store(key, value string) (Store, error) {
	if i, err := s.FindKV(key); err != nil {
		s.Data = append(s.Data, KeyValue{
			Key:   key,
			Value: value,
		})
	} else {
		s.Data[i] = KeyValue{
			Key:      key,
			Value:    value,
			Revision: s.Data[i].Revision + 1,
		}
	}
	err := s.writedb()
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s Store) Dump() Store {
	return s
}

func (s Store) RunCmd() error {
	if len(os.Args) < 2 {
		fmt.Println(s.Dump())
		return nil
	}
	key := os.Args[1]
	
	var value string
	if len(os.Args) >= 3 {
		value = os.Args[2]
		fmt.Println(value)
		_, err := s.Store(key, value)
		if err != nil {
			return err
		}
		return nil
	}

	if value = s.Retrieve(key); value != "" {
		fmt.Println(value)		
	}
	return nil
}