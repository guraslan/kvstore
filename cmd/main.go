package main

import (
	"fmt"
	"os"
	"github.com/guraslan/kvstore"
)

func main() {
   s := kvstore.Open()
   if len(os.Args) < 2 {
      fmt.Println(s.Dump())
	  os.Exit(0)
   }
   key := os.Args[1]
   
   var value string
   if len(os.Args) == 3 {
	 value = os.Args[2]
	 s.Store(key, value)
	 os.Exit(0)
   }

   value = s.Retrieve(key)
   if value == "" {
      os.Exit(0)
   }
   fmt.Println(value)
}
