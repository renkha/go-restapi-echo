package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/renkha/go-restapi/helper"
	"github.com/renkha/go-restapi/src/usecase"
	re "github.com/renkha/go-restapi/src/usecase/user"
)

type userDelivery struct {
	userUsecase usecase.UserUsecase
}

func NewDelivery(userUsecase usecase.UserUsecase) *userDelivery {
	return &userDelivery{userUsecase}
}

func (d *userDelivery) UserRegistration(c echo.Context) error {
	req := new(re.UserRequest)

	if err := c.Bind(req); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(req); err != nil {
		// errors := helper.ErrorFormatter(err)
		// errMessage := helper.M{"errors": errors}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	existEmail := d.userUsecase.CheckExistEmail(*req)
	if existEmail != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", existEmail.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := d.userUsecase.CreateUser(*req)
	if err != nil {
		// errors := helper.ErrorFormatter(err)
		// errMessage := helper.M{"errors": errors}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "failed", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	userData := re.UserResponseFormatter(newUser)

	response := helper.ResponseFormatter(http.StatusOK, "success", "success user", userData)

	return c.JSON(http.StatusOK, response)
}
