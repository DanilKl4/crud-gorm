package main

import (
	"flag"
	"github.com/DanilKl4/crud-gorm/internal/app"
	"log"
)

var (
	f = flag.String("p", "5000", "port on the server listening on")
)

func main() {
	flag.Parse()

	err := app.Run(*f)
	if err != nil {
		log.Fatal(err)
	}
}
