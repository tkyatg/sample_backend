package http

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	usercommand "github.com/tkyatg/rental_redy_backend/services/user/commands/userCommand"
	"github.com/tkyatg/rental_redy_backend/services/user/domain"
	userquery "github.com/tkyatg/rental_redy_backend/services/user/queries/userQuery"
)

func NewEchoServer(dbConnection *sql.DB) *echo.Echo {
	echo := echo.New()
	echo.Use(middleware.Logger())

	// userQueryService
	userQueryDataAccessor := userquery.NewDataAccessor(dbConnection)
	userQueryUsecase := userquery.NewUsecase(userQueryDataAccessor)
	userquery.NewServer(echo, userQueryUsecase)

	// userCommandService
	userCommandDataAccessor := domain.NewUserDataAccessor(dbConnection)
	userCommandRepository := domain.NewUserRepository(userCommandDataAccessor)
	userCommandUsecase := usercommand.NewUsecase(userCommandRepository)
	usercommand.NewServer(echo, userCommandUsecase)

	return echo
}
