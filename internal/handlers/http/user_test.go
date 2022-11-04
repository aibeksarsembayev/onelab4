package httpdelivery_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	httpdelivery "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/handlers/http"
	_userRepo "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/repository"
	_userUsecase "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/usecases"
	"github.com/jmoiron/sqlx"

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

	// load configs
	// conf, err := config.LoadConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(conf)
	// }

	// dbpool, err := postgres.InitPostgresDBConn(&conf)
	// if err != nil {
	// 	log.Fatalf("database: %v", err)
	// }
	// defer dbpool.Close()

	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	userRepo := _userRepo.NewDBUserRepository(sqlxDB) // pass DB conn
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
