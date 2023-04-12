package main

import (
	"assigment10/database"
	"assigment10/router"
)

func main() {
	database.StartDB()

	router.StartApp().Run(":8080")
}
