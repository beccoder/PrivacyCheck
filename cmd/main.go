package main

import (
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"privacy-check/configs/env"
	"privacy-check/configs/pg"
	"privacy-check/database"
	"privacy-check/internal/handler"
	"privacy-check/internal/repository"
	"privacy-check/internal/service"
)

// @title Privacy Check Server API
// @version 1.0
// @description Privacy Check Server Api in Go

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	projectEnv := env.ProjectEnv()

	pgConfig := pg.NewConfigEmpty()
	{
		pgConfig.SetHost(projectEnv.PgHost).
			SetPort(projectEnv.PgPort).
			SetDbname(projectEnv.PgDB).
			SetUser(projectEnv.PgUser).
			SetPassword(projectEnv.PgPassword)
	}

	db, err := database.Connect(pgConfig)
	{
		if err != nil {
			log.Panicln(err)
		}
	}

	router := echo.New()
	{
		setMiddlewares(router)
		createHandler(router, db, &projectEnv)
		runHTTPServerOnAddr(router, projectEnv.HttpPort)
	}
}

func runHTTPServerOnAddr(handler *echo.Echo, port int) {
	url := strconv.FormatInt(int64(port), 10)
	{
		log.Panicln(handler.Start(":" + url))
	}
}

func setMiddlewares(router *echo.Echo) {
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(middleware.RequestID())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
}

func createHandler(router *echo.Echo, db *sqlx.DB, config *env.EnvProject) {
	group := router.Group("/api/v1")
	{
		handler.NewHandler(group, service.NewService(repository.NewRepository(db), config))
	}
}
