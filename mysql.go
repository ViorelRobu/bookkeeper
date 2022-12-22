package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type BooksCollection struct {
	Result []Book
}

func getAllBooks() BooksCollection {
	results, queryErr := DB.Query("SELECT * FROM books")

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

	return AllBooks
}
