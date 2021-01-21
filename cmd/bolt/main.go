package main

import (
	"os"
	"github.com/boltdb/bolt/cmd/bolt/mainhelper"
	"fmt"
)
func main() {
m := mainhelper.NewMain()
if err := m.Run(os.Args[1:]...); err == mainhelper.ErrUsage {
os.Exit(2)
} else if err != nil {
fmt.Println(err.Error())
os.Exit(1)
}
}
