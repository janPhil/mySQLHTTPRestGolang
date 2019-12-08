package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/janPhil/mySQLHTTPRestGolang/server"
	"log"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("Couldn't start the server: %v", err)
	}
	s.StartServer()

}
