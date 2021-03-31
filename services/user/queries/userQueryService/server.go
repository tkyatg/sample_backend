package userqueryservice

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tkyatg/rental_redy_backend/shared"
)

type (
	server struct {
		uc Usecase
	}
)

func NewServer(e *echo.Echo, us Usecase) {
	server := &server{
		us,
	}
	e.GET("/users", server.GetUserList)
	e.GET("/users/:uuid", server.GetUserByID)
}

func (s *server) GetUserList(etx echo.Context) error {
	users, err := s.uc.getUserList()
	if err != nil {
		return etx.JSON(http.StatusNotFound, shared.NewResponseError(err.Error()))
	}
	return etx.JSON(http.StatusOK, users)
}

func (s *server) GetUserByID(etx echo.Context) error {
	userUUID := etx.Param("uuid")
	user, err := s.uc.getUserByID(userUUID)
	if err != nil {
		return etx.JSON(http.StatusNotFound, shared.NewResponseError(err.Error()))
	}
	return etx.JSON(http.StatusOK, user)
}
