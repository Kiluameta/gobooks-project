package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Kiluameta/gobooks-project/internal/cli"
	"github.com/Kiluameta/gobooks-project/internal/service"
	"github.com/Kiluameta/gobooks-project/internal/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./book.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(bookService)

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		bookCLI := cli.NewBookCLI(bookService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBooksByID)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}
