package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ebenne01/bct-go-assessment/model"
	"github.com/labstack/echo/v4"
)

const (
	_expectedStatusCode = "Expected %d status code"
	_unexpectedError    = "Unexpected error: %s"
)

type mockUser struct {
	err error
}

func (mock *mockUser) GetAllUsers() ([]model.User, error) {
	return []model.User{}, mock.err
}

func (mock *mockUser) Create() (model.User, error) {
	return model.User{}, mock.err
}

func (mock *mockUser) Update(id int) error {
	return mock.err
}

func (mock *mockUser) Delete(id int) error {
	return mock.err
}

type mockUserFactory struct {
	err error
}

func (m *mockUserFactory) createUser() model.UserModel {
	return &mockUser{err: m.err}
}

func TestGetAllUsers(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	c := e.NewContext(req, rec)

	if err := GetAllUsers(c); err != nil {
		t.Errorf(_unexpectedError, err.Error())
	}
}

func TestGetAllUsersReturns500Error(t *testing.T) {
	mock := mockUserFactory{err: errors.New("Error getting all users")}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	c := e.NewContext(req, rec)

	if err := GetAllUsers(c); err == nil {
		t.Errorf(_expectedStatusCode, http.StatusInternalServerError)
	} else {
		if errCode := err.(*echo.HTTPError).Code; errCode != http.StatusInternalServerError {
			t.Errorf(_expectedStatusCode, http.StatusInternalServerError)
		}
	}
}

func TestCreateUser(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)

	c := e.NewContext(req, rec)

	if err := CreateUser(c); err != nil {
		t.Errorf(_expectedStatusCode, http.StatusCreated)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf(_expectedStatusCode, http.StatusCreated)
	}
}

func TestCreateUserHandlesDuplicateUserName(t *testing.T) {
	mock := mockUserFactory{err: model.ErrDuplicateUserName}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)

	c := e.NewContext(req, rec)

	if err := CreateUser(c); err == nil {
		t.Errorf(_expectedStatusCode, http.StatusConflict)
	} else {
		if errCode := err.(*echo.HTTPError).Code; errCode != http.StatusConflict {
			t.Errorf(_expectedStatusCode, http.StatusConflict)
		}
	}
}

func TestCreateUserHandlesUnknownError(t *testing.T) {
	mock := mockUserFactory{err: errors.New("some other error")}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)

	c := e.NewContext(req, rec)

	err := CreateUser(c)
	if err == nil {
		t.Errorf(_expectedStatusCode, http.StatusBadRequest)
	} else {
		if errCode := err.(*echo.HTTPError).Code; errCode != http.StatusBadRequest {
			t.Errorf(_expectedStatusCode, http.StatusBadRequest)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/users/123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := UpdateUser(c)
	if err != nil {
		t.Errorf(_unexpectedError, err.Error())
	}

	if rec.Code != http.StatusNoContent {
		t.Errorf(_expectedStatusCode, http.StatusNoContent)
	}
}

func TestUpdateUserHandlesBadId(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/users/a123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("a123")

	err := UpdateUser(c)
	if err == nil {
		t.Errorf(_expectedStatusCode, http.StatusBadRequest)
	} else {

		if !strings.Contains(err.Error(), BadIdErrorMsg) {
			t.Errorf(_unexpectedError, err.Error())
		}
	}
}

func TestUpdateUserHandlesError(t *testing.T) {
	mock := mockUserFactory{err: errors.New("foo")}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/users/123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := UpdateUser(c)
	if err == nil {
		t.Errorf(_expectedStatusCode, http.StatusBadRequest)
	} else {
		if errCode := err.(*echo.HTTPError).Code; errCode != http.StatusBadRequest {
			t.Errorf(_unexpectedError, err.Error())
		}
	}
}

func TestDeleteUser(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/users/123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := DeleteUser(c)
	if err != nil {
		t.Errorf(_unexpectedError, err.Error())
	}

	if rec.Code != http.StatusNoContent {
		t.Errorf(_expectedStatusCode, http.StatusNoContent)
	}
}

func TestDeleteUserHandlesBadId(t *testing.T) {
	mock := mockUserFactory{err: nil}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/users/a123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("a123")

	err := DeleteUser(c)
	if err == nil {
		t.Errorf(_expectedStatusCode, http.StatusBadRequest)
	} else {

		if !strings.Contains(err.Error(), BadIdErrorMsg) {
			t.Errorf(_unexpectedError, err.Error())
		}
	}
}

func TestDeleteUserHandlesError(t *testing.T) {
	mock := mockUserFactory{err: errors.New("foo")}
	userFactory = &mock

	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/users/123", nil)

	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := DeleteUser(c)
	if err == nil {
		t.Errorf(_expectedStatusCode, http.StatusBadRequest)
	} else {
		if errCode := err.(*echo.HTTPError).Code; errCode != http.StatusBadRequest {
			t.Errorf(_unexpectedError, err.Error())
		}
	}
}
