package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type BooksCollection struct {
	Result []Book
}

func getAllBooks() BooksCollection {
	db, err := sql.Open("mysql", "root:password@tcp(database:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	results, queryErr := db.Query("SELECT * FROM books")

	if queryErr != nil {
		panic(queryErr.Error())
	}

	var AllBooks BooksCollection

	for results.Next() {
		var book Book
		scanErr := results.Scan(&book.Id, &book.Name, &book.Author, &book.ISBN)

		if scanErr != nil {
			panic(scanErr.Error())
		}

		AllBooks.Result = append(AllBooks.Result, book)
	}

	defer db.Close()

	return AllBooks
}
