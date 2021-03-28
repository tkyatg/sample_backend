package userqueryservice

import (
	"net/http"

	"github.com/labstack/echo"
)

type server struct {
	uc Usecase
}

func NewServer(e *echo.Echo, us Usecase) {
	handler := &server{
		us,
	}
	e.GET("/", handler.Test)
}

func (s *server) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "sample")
}
