package domain

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

type (
	repositoryTestHelper struct {
		t    *testing.T
		etx  *echo.Echo
		ctrl *gomock.Controller
		da   *MockUserDataAccessor
		repo UserRepository
	}
)

func newRepositoryTestHelper(t *testing.T) *repositoryTestHelper {
	ctrl := gomock.NewController(t)
	etx := echo.New()
	da := NewMockUserDataAccessor(ctrl)
	repo := NewUserRepository(da)

	return &repositoryTestHelper{t, etx, ctrl, da, repo}
}

func TestRepositoryCreate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newRepositoryTestHelper(t)
		defer h.ctrl.Finish()

		h.da.EXPECT().create(&CreateUserAttributes{}).Return(nil)

		if actualErr := h.repo.Create(&CreateUserAttributes{}); actualErr != nil {
			t.Fatal(actualErr)
		}
	})
	t.Run("異常系", func(t *testing.T) {
		t.Parallel()
		h := newRepositoryTestHelper(t)
		defer h.ctrl.Finish()
		err := errors.New("error")

		h.da.EXPECT().create(&CreateUserAttributes{}).Return(err)

		if actualErr := h.repo.Create(&CreateUserAttributes{}); actualErr != err {
			t.Fatal(actualErr)
		}

	})
}
