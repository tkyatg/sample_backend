package usercommand

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	server struct {
		uc Usecase
	}
	createRequest struct {
		Email    string
		Password string
		Gender   string
	}
)

func NewServer(e *echo.Echo, us Usecase) {
	server := &server{
		us,
	}
	e.POST("/users/new", server.Create)
}

func (s *server) Create(e echo.Context) error {
	req := createRequest{
		Email:    e.FormValue("email"),
		Password: e.FormValue("password"),
		Gender:   e.FormValue("gender"),
	}
	err := s.uc.Create(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return e.JSON(http.StatusCreated, "ok")
}
