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
		Create(context.Context, *models.Book) error
	}
}

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@/bookstore?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
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
	app.POST("/books", app.booksCreate)
}

func (app *App) booksCreate(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	var book models.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(400, "invalid data")
		return err
	}

	err := app.books.Create(ctx, &book)
	if err != nil {
		c.JSON(500, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(200, book)
}

func (app *App) booksIndex(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	defer cancel()

	bks, err := app.books.All(ctx)
	if err != nil {
		c.JSON(500, map[string]string{"error": "internal server error"})
		return err
	}

	return c.JSON(200, bks)
}
