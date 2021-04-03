package usercommand

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/tkyatg/rental_redy_backend/services/user/domain"
)

type (
	usecaseTestHelper struct {
		t       *testing.T
		etx     *echo.Echo
		ctrl    *gomock.Controller
		usecase Usecase
		repo    *domain.MockUserRepository
	}
)

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	etx := echo.New()
	repo := domain.NewMockUserRepository(ctrl)
	usecase := NewUsecase(repo)

	return &usecaseTestHelper{t, etx, ctrl, usecase, repo}
}

func TestUsecaseCreate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newUsecaseTestHelper(t)
		defer h.ctrl.Finish()
		req := createRequest{
			Email:    "test@gmail.com",
			Password: "password",
			Gender:   "男",
		}
		h.repo.EXPECT().Create(gomock.Any()).Return(nil)

		if actualErr := h.usecase.Create(req); actualErr != nil {
			t.Fatal(actualErr)
		}
	})
	t.Run("異常系", func(t *testing.T) {
		t.Parallel()
		h := newUsecaseTestHelper(t)
		defer h.ctrl.Finish()

		err := errors.New("error")
		req := createRequest{
			Email:    "test@gmail.com",
			Password: "password",
			Gender:   "男",
		}
		h.repo.EXPECT().Create(gomock.Any()).Return(err)

		if actualErr := h.usecase.Create(req); actualErr != err {
			t.Fatal(actualErr)
		}
	})
}
