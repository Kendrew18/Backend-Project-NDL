package main

import (
	"Backend-Project-NDL/db"
	"Backend-Project-NDL/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
