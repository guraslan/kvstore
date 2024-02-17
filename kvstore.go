package kvstore

import (
	"fmt"
	"os"
)

var storefile = "global.db"

type kv struct {
	key      string
	value    string
	revision int
}

type store []kv

func Open() store {
	s := make(store, 0)
	// TODO: Unless it's permission denied accessing/creating
	// is it safe to ignore the error?
	s, _ = s.readdb()
	return s
}

func (s store) readdb() (store, error) {
	f, err := os.OpenFile(storefile, os.O_RDONLY|os.O_CREATE, 0700)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for {
		var key, value string
		var revision int
		_, err := fmt.Fscanf(f, "%v\n%v\n%v\n", &key, &value, &revision)
		if err != nil {
			break
		}

		pair := kv{
			key:      key,
			value:    value,
			revision: revision,
		}
		s = append(s, pair)
	}
	return s, nil
}

func (s store) writedb() error {
	f, err := os.OpenFile(storefile, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf string
	for _, pair := range s {
		buf += fmt.Sprintf("%v\n%v\n%v\n", pair.key, pair.value, pair.revision)
	}
	f.Write([]byte(buf))
	return nil
}

func (s store) Retrieve(key string) string {
	i, err := s.findKV(key)
	if err != nil {
		return ""
	}
	return s[i].value
}

func (s store) findKV(key string) (int, error) {
	for i, r := range s {
		if r.key == key {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

func (s store) Store(key, value string) {
	if i, err := s.findKV(key); err != nil {
		s = append(s, kv{
			key:   key,
			value: value,
		})
	} else {
		s[i] = kv{
			key:      key,
			value:    value,
			revision: s[i].revision + 1,
		}
	}
	err := s.writedb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (s store) Dump() store {
	return s
}
