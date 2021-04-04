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
		Email    string `json:"email"`
		Password string `json:"password"`
		Gender   string `json:"gender"`
	}
)

func NewServer(e *echo.Echo, us Usecase) {
	server := &server{
		us,
	}
	e.POST("/users/new", server.Create)
}

func (s *server) Create(e echo.Context) error {
	var req createRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err := s.uc.Create(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return e.JSON(http.StatusCreated, "ok")
}
