package helper

import "github.com/labstack/echo/v4"

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlerware []echo.MiddlewareFunc
}

type Handler interface {
	Route() []Route
}
