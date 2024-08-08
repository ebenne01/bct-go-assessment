package controller

import (
	"net/http"
	"strconv"

	"github.com/ebenne01/bct-go-assessment/model"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users, err := model.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "oops, something went wrong")
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	newUser, err := model.Create(*user)
	if err != nil {
		if err == model.ErrDuplicateUserName {
			return echo.NewHTTPError(http.StatusConflict, err)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
	}

	return c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c echo.Context) error {

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameter value.  'ID' must  be numeric")
	}

	err = model.Update(id, *user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid parameter value.  'ID' must  be numeric")
	}

	err = model.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusNoContent)
}
