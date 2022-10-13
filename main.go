package main

import (
	"context"
	"net/http"
	"time"

	"bookstore.splindid/models"
	"bookstore.splindid/pkg/ratelimit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	app.Use(middleware.Logger())

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0, // use default DB
	})
	redis_rate.PerMinute(10)
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store:   ratelimit.NewRedisRateLimiter(rdb, redis_rate.PerMinute(10)),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	app.Use(middleware.RateLimiterWithConfig(config))

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
