package httpdelivery

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
	"github.com/labstack/echo/v4"
)

// ResponceError
type ResponceError struct {
	Message string `json:"message"`
}

// Create user ...
func (uh *UserHandler) Create(c echo.Context) error {
	user := new(domain.UserRequestDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// add user to db
	us := domain.User{
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Status:    true,
		CreatedAt: time.Now(),
	}
	err := uh.UserUsecase.Create(c.Request().Context(), &us)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return c.JSON(http.StatusBadRequest, ResponceError{Message: "user already exists"})
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusCreated, user)
}

// Get user by ID ...
func (uh *UserHandler) GetByID(c echo.Context) error {
	// get user id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// get user from db
	user, err := uh.UserUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		fmt.Println(err.Error() == "no record")
		if strings.Contains(err.Error(), "no rows in result set") {
			return c.JSON(http.StatusNotFound, ResponceError{Message: "no user with given id"})
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusFound, user)
}

// Get all users ...
func (uh *UserHandler) GetAll(c echo.Context) error {
	users, err := uh.UserUsecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, ResponceError{Message: "no users"})
	}

	return c.JSON(http.StatusOK, users)
}

// Update user by id ...
func (uh *UserHandler) Update(c echo.Context) error {
	// get id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// user info bind from form
	user := new(domain.UserRequestDTO)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update user in db
	us := domain.User{
		ID:      id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}
	err = uh.UserUsecase.Update(context.Background(), &us)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return c.JSON(http.StatusNotFound, ResponceError{Message: "no user with given id"})
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, user)
}

// Delete user by id ...
func (uh *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = uh.UserUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				return c.JSON(http.StatusNotFound, ResponceError{Message: "no user with given id"})
			} else {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
	}
	return c.JSON(http.StatusOK, ResponceError{Message: "succcessfully deleted"})
}
