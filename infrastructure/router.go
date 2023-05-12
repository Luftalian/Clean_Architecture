package infrastructure

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srinathgs/mysqlstore"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Luftalian/Clean_Architecture/interfaces/controller"
)

var Router *echo.Echo

func init() {
	dbHandler := NewDbHandler()
	store, err := mysqlstore.NewMySQLStoreFromConnection(dbHandler.Conn.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	uuidHandler := new(UUIDHandler)

	userController := controller.NewUserController(dbHandler)

	// e := e.Group("/api")
	e.POST("/users", func(c echo.Context) error {
		userController.Create(c, uuidHandler)
		return nil
	})
	e.GET("/users", func(c echo.Context) error {
		userController.Index(c)
		return nil
	})
	e.GET("/users/:id", func(c echo.Context) error {
		userController.Show(c, uuidHandler)
		return nil
	})

	Router = e
}
