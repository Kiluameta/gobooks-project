package cli

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Kiluameta/gobooks-project/internal/service"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usege: books <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book_title>")
			return
		}
		bookName := os.Args[2]
		cli.SearchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
			return
		}
		bookIDs := os.Args[2:]
		cli.SimulateReading(bookIDs)
	}
}

func (cli *BookCLI) SearchBooks(name string) {
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}

	fmt.Printf("%d books found\n", len(books))
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n",
			book.ID, book.Title, book.Author, book.Genre,
		)
	}
}

func (cli *BookCLI) SimulateReading(bookIDsStr []string) {
	var bookIDs []int
	for _, idStr := range bookIDsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid book ID:", idStr)
			continue
		}
		bookIDs = append(bookIDs, id)
	}

	responses := cli.service.SimulateMultiplereadings(bookIDs, 2*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}
