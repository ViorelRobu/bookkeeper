package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Name string
	User string
	Pass string
}

func main() {
	fmt.Println("Starting web server on port :8080...")
	fmt.Println("Web server started - ok")

	db := Database{}
	db.getCredsFromEnv()
	HandleRoutes()

	fmt.Println(db)

	serverStartError := http.ListenAndServe(":8080", nil)

	if serverStartError != nil {
		log.Fatal(serverStartError)
	}
}

func (db *Database) getCredsFromEnv() {
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USER")
	db.Pass = os.Getenv("DB_PASSWORD")
}
