package main

import (
	"fmt"
	"os"
	"github.com/guraslan/kvstore"
)

func main() {
   if len(os.Args) < 2 {
      fmt.Println(kvstore.Dump())
	  os.Exit(0)
   }
   key := os.Args[1]
   
   var value string
   if len(os.Args) == 3 {
	 value = os.Args[2]
	 kvstore.Store(key, value)
	 os.Exit(0)
   }

   value = kvstore.Retrieve(key)
   if value == "" {
      os.Exit(0)
   }
   fmt.Println(value)
}
