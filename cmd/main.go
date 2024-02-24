package main

import (
	"fmt"
	"os"

	"github.com/guraslan/kvstore"
)

func main() {
	s, err := kvstore.OpenDefault()
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	s.RunCmd()
}
