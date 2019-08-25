package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/janPhil/mySQLHTTPRestGolang/server"
)

func main() {

	s := server.NewServer()
	s.StartServer()

}
