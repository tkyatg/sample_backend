package userqueryservice

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	server struct {
		uc Usecase
	}
	ResponseError struct {
		Message string `json:"message"`
	}
	getUserByIDResult struct {
		UserUUID         string `gorm:"primary_key"`
		DisplayName      string `gorm:"column:display_name"`
		BirthDay         string `gorm:"column:birthday"`
		Gender           string `gorm:"column:gender"`
		ImageURL         string `gorm:"column:image_url"`
		FreeTime         string `gorm:"column:free_time"`
		SelfIntroduction string `gorm:"column:self_introduction"`
		CreatedAt        string `gorm:"column:created_at"`
		UpdatedAt        string `gorm:"column:updated_at"`
	}
	getUserListResult struct {
		UserUUID         string `gorm:"primary_key"`
		DisplayName      string `gorm:"column:display_name"`
		BirthDay         string `gorm:"column:birthday"`
		Gender           string `gorm:"column:gender"`
		ImageURL         string `gorm:"column:image_url"`
		FreeTime         string `gorm:"column:free_time"`
		SelfIntroduction string `gorm:"column:self_introduction"`
		CreatedAt        string `gorm:"column:created_at"`
		UpdatedAt        string `gorm:"column:updated_at"`
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
		return etx.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return etx.JSON(http.StatusOK, users)
}

func (s *server) GetUserByID(etx echo.Context) error {
	userUUID := etx.Param("uuid")
	user, err := s.uc.getUserByID(userUUID)
	if err != nil {
		return etx.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return etx.JSON(http.StatusOK, user)
}
