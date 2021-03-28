package http

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	userqueryservice "github.com/tkyatg/rental_redy_backend/services/user/queries/userQueryService"
)

func NewEchoServer(dbConnection *gorm.DB) *echo.Echo {
	echo := echo.New()

	// userQueryService
	userQueryServiceDa := userqueryservice.NewDataAccessor(dbConnection)
	userQueryServiceUc := userqueryservice.NewUsecase(userQueryServiceDa)
	userqueryservice.NewServer(echo, userQueryServiceUc)

	return echo
}
