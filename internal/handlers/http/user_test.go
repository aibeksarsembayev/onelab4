package httpdelivery_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	httpdelivery "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/handlers/http"
	_userRepo "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/repository"
	"github.com/aibeksarsembayev/onelab/tasks/lab4/internal/repository/postgres"
	_userUsecase "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/usecases"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"name":"bender","surname":"unknown","email":"r3000@m.com"}
`
)

func TestCreate(t *testing.T) {
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("name", "bender")
	f.Set("surname", "unknown")
	f.Set("email", "r3000@m.com")

	req := httptest.NewRequest(http.MethodPost, "/user/create", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// initialize repos. 	// set context timeout and initialize usecases
	dbpool, err := postgres.InitPostgresDBConn(&conf)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer dbpool.Close()
	userRepo := _userRepo.NewDBUserRepository(dbpool) // pass DB conn
	timeoutContext := 2 * time.Second
	uUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	// handler
	h := &httpdelivery.UserHandler{
		UserUsecase: uUsecase,
	}

	// Assertions
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
