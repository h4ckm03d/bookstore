package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"bookstore.splindid/models"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	// Replace the reference to models.BookModel with an interface
	// describing its methods instead. All the other code remains exactly
	// the same.
	books interface {
		All(context.Context) ([]models.Book, error)
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
	fmt.Println("starting server at :3000")
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	bks, err := env.books.All(ctx)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(bks)
}
