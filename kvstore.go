package kvstore

import (
	"fmt"
	"os"
)

var storefile = "global.db"

var store = map[string]string{}

func readdb() error {
	f, err := os.OpenFile(storefile, os.O_RDONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		var k, v string
		_, err := fmt.Fscanf(f, "%s\n%s\n", &k, &v)
		if err != nil {
			break
		}
		store[k] = v
	}
	return nil
}

func writedb() error {
	f, err := os.OpenFile(storefile, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf string
	for k, v := range store {
		buf = buf + k + "\n" + v + "\n"
	}
	f.Write([]byte(buf))
	return nil
}

func Retrieve(key string) string {
	readdb()
	return store[key]
}

func Store(key, value string) {
	readdb()
	store[key] = value
	err := writedb()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Dump() map[string]string {
	readdb()
	return store
}
