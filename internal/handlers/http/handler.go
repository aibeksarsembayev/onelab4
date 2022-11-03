package httpdelivery

import (
	"net/http"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
	"github.com/labstack/echo/v4"
)

// UseHandler represent the httphandler for user
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler ...
func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! It is simple CRUD template on Echo with Docker and GitLab CI/CD")
	})
	// user routes
	user := e.Group("/user")
	user.POST("/create", handler.Create)
	user.GET("/:id", handler.GetByID)
	user.GET("/all", handler.GetAll)
	user.PUT("/:id", handler.Update)
	user.DELETE("/delete", handler.Delete)
}
