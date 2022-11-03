// создать template crud используя фреймворк echo, (4 роута, разные методы http, разные способы передачи данных - by query, by body, etc.)
// завернуть в докер контейнер, и запустить локально
// проверить роуты апи через постман либо через запросы с другого го инстанса (net/http)
// загрузить лабку на гитхаб/гитлаб
package main

import (
	"fmt"
	"log"
	"time"

	config "github.com/aibeksarsembayev/onelab/tasks/lab4/config"
	_userHttpDelivery "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/handlers/http"
	_userRepo "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/repository"
	"github.com/aibeksarsembayev/onelab/tasks/lab4/internal/repository/postgres"
	_userUsecase "github.com/aibeksarsembayev/onelab/tasks/lab4/internal/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load configs
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conf)
	}

	// create pool of connection for DB
	dbpool, err := postgres.InitPostgresDBConn(&conf)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer dbpool.Close()

	// initialize echo server
	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())

	// initialize repos
	userRepo := _userRepo.NewDBUserRepository(dbpool)

	// set context timeout
	timeoutContext := time.Duration(conf.Context.Timeout) * time.Second

	// initialize usecases
	uUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	// initialzie handlers
	_userHttpDelivery.NewUserHandler(e, uUsecase)

	// start echo server
	e.Logger.Fatal(e.Start(conf.Server.Address))
}
