package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {
	fmt.Println("Starting web server on port :8080...")
	fmt.Println("Web server started - ok")

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	conn, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(database:3306)/"+dbName)
	DB = conn
	if err != nil {
		panic(err.Error())
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	HandleRoutes()

	serverStartError := http.ListenAndServe(":8080", nil)

	if serverStartError != nil {
		log.Fatal(serverStartError)
	}
}
