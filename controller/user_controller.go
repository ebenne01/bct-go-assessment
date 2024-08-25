package controller

import (
	"net/http"
	"strconv"

	"github.com/ebenne01/bct-go-assessment/model"
	"github.com/labstack/echo/v4"
)

const (
	BadIdErrorMsg = "ID must be numeric"
)

type UserFactory interface {
	createUser() model.UserModel
}

type concreteUserFactory struct{}

func (cu *concreteUserFactory) createUser() model.UserModel {
	user := new(model.User)
	return user
}

var userFactory UserFactory = &concreteUserFactory{}

func GetAllUsers(c echo.Context) error {
	user := userFactory.createUser()

	users, err := user.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "oops, something went wrong")
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := userFactory.createUser()

	if err := c.Bind(user); err != nil {
		return err
	}

	newUser, err := user.Create()
	if err != nil {
		if err == model.ErrDuplicateUserName {
			return echo.NewHTTPError(http.StatusConflict, err)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	user := userFactory.createUser()

	if err := c.Bind(user); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, BadIdErrorMsg)
	}

	err = user.Update(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteUser(c echo.Context) error {
	user := userFactory.createUser()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, BadIdErrorMsg)
	}

	err = user.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)
}
