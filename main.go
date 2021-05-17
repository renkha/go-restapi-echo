package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/renkha/go-restapi/src"
)

//validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func main() {
	//load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("fail to load .env file")
	}

	//new echo
	e := echo.New()

	//validator
	e.Validator = &CustomValidator{validator: validator.New()}

	//logger by echo
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	//remove slash(/) at the end of endpoint
	e.Pre(middleware.RemoveTrailingSlash())

	//routes
	src.DefineApiRoutes(e)

	//port
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
