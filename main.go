package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore.majoo/models"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	// Replace the reference to models.BookModel with an interface
	// describing its methods instead. All the other code remains exactly
	// the same.
	books interface {
		All() (context.Context, []models.Book, error)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", env.booksIndex)

	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := env.books.All()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
