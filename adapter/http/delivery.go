package http

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	userquery "github.com/tkyatg/rental_redy_backend/services/user/queries/userQuery"
)

func NewEchoServer(dbConnection *gorm.DB) *echo.Echo {
	echo := echo.New()

	// userQueryService
	userQueryDataAccessor := userquery.NewDataAccessor(dbConnection)
	userQueryUsecase := userquery.NewUsecase(userQueryDataAccessor)
	userquery.NewServer(echo, userQueryUsecase)

	return echo
}
