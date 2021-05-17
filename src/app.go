package src

import (
	"github.com/labstack/echo/v4"
	"github.com/renkha/go-restapi/config"
	"github.com/renkha/go-restapi/helper"
	"github.com/renkha/go-restapi/src/delivery"
	"github.com/renkha/go-restapi/src/model"
	"github.com/renkha/go-restapi/src/repository"
	"github.com/renkha/go-restapi/src/usecase"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := config.GetDbInstance()
	db.AutoMigrate(model.User{})

	userRepository := repository.NewRepository(db)
	userUsecase := usecase.NewUsecase(userRepository)
	userDelivery := delivery.NewDelivery(userUsecase)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userDelivery.UserRegistration,
		},
	}
}

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handler{
		UserRoutes{},
	}
	var routes []helper.Route

	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api/v1")

	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			api.POST(route.Path, route.Handler, route.Middlerware...)
		case echo.GET:
			api.GET(route.Path, route.Handler, route.Middlerware...)
		case echo.PUT:
			api.PUT(route.Path, route.Handler, route.Middlerware...)
		case echo.DELETE:
			api.DELETE(route.Path, route.Handler, route.Middlerware...)
		}
	}
}
