package userquery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tkyatg/rental_redy_backend/shared"
)

type (
	server struct {
		uc Usecase
	}
	getUserByIDRequest struct {
		userUUID string
	}
)

func NewServer(e *echo.Echo, us Usecase) {
	server := &server{
		us,
	}
	e.GET("/users", server.GetUserList)
	e.GET("/users/:uuid", server.GetUserByID)
}

func (s *server) GetUserList(e echo.Context) error {
	users, err := s.uc.getUserList()
	if err != nil {
		return e.JSON(http.StatusNotFound, shared.NewResponseError(err.Error()))
	}
	return e.JSON(http.StatusOK, users)
}

func (s *server) GetUserByID(e echo.Context) error {
	req := getUserByIDRequest{userUUID: e.Param("uuid")}
	user, err := s.uc.getUserByID(req)
	if err != nil {
		return e.JSON(http.StatusNotFound, shared.NewResponseError(err.Error()))
	}
	return e.JSON(http.StatusOK, user)
}
