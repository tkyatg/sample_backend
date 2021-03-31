package http

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	authenticationcommand "github.com/tkyatg/rental_redy_backend/services/auth/commands/authenticationCommand"
	userqueryservice "github.com/tkyatg/rental_redy_backend/services/user/queries/userQueryService"
)

func NewEchoServer(dbConnection *gorm.DB) *echo.Echo {
	echo := echo.New()

	// userQueryService
	userQueryServiceDa := userqueryservice.NewDataAccessor(dbConnection)
	userQueryServiceUc := userqueryservice.NewUsecase(userQueryServiceDa)
	userqueryservice.NewServer(echo, userQueryServiceUc)

	// authenticationCommand
	authenticationCommandUc := authenticationcommand.NewUsecase()
	authenticationcommand.NewServer(echo, authenticationCommandUc)

	return echo
}
