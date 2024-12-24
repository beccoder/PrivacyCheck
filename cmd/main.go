package main

import (
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	loader "privacy-check/configs/env"
	"privacy-check/configs/pg"
	"privacy-check/database"
	"privacy-check/internal/handler"
	"privacy-check/internal/repository"
	"privacy-check/internal/service"
)

func main() {
	projectEnv := loader.ProjectEnv()

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
		createHandler(router, db)
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

func createHandler(router *echo.Echo, db *sqlx.DB) {
	group := router.Group("/api/v1")
	{
		handler.NewHandler(group, service.NewService(repository.NewRepository(db)))
	}
}
