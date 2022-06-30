package main

import (
	"context"
	"time"

	"bookstore.splindid/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type App struct {
	*echo.Echo

	books interface {
		All(context.Context) ([]models.Book, error)
	}
}

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@/bookstore"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	app := &App{
		books: models.BookModel{DB: db},
		Echo:  echo.New(),
	}
	app.SetupRoutes()
	app.Logger.Fatal(app.Start(":3000"))
}

func (app *App) SetupRoutes() {
	app.GET("/books", app.booksIndex)
}

func (app *App) booksIndex(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	bks, err := app.books.All(ctx)
	if err != nil {
		c.JSON(500, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(200, bks)
}
