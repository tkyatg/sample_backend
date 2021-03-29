package userqueryservice

import (
	"errors"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

type (
	usecaseTestHelper struct {
		t            *testing.T
		etx          *echo.Echo
		ctrl         *gomock.Controller
		usecase      Usecase
		dataAccessor *MockDataAccessor
	}
)

func newUsecaseTestHelper(t *testing.T) *usecaseTestHelper {
	ctrl := gomock.NewController(t)
	etx := echo.New()
	dataAccessor := NewMockDataAccessor(ctrl)
	usecase := NewUsecase(dataAccessor)

	return &usecaseTestHelper{t, etx, ctrl, usecase, dataAccessor}
}

func TestUsecaseGetUsers(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newUsecaseTestHelper(t)
		defer h.ctrl.Finish()

		users := []getUserListResult{
			{
				UserUUID:         "UserUUID-01",
				DisplayName:      "DisplayName-01",
				BirthDay:         "BirthDay-01",
				Gender:           "Gender-01",
				ImageURL:         "ImageURL-01",
				FreeTime:         "FreeTime-01",
				SelfIntroduction: "SelfIntroduction-01",
				CreatedAt:        "CreatedAt-01",
				UpdatedAt:        "UpdatedAt-01",
			},
			{
				UserUUID:         "UserUUID-02",
				DisplayName:      "DisplayName-02",
				BirthDay:         "BirthDay-02",
				Gender:           "Gender-02",
				ImageURL:         "ImageURL-02",
				FreeTime:         "FreeTime-02",
				SelfIntroduction: "SelfIntroduction-02",
				CreatedAt:        "CreatedAt-02",
				UpdatedAt:        "UpdatedAt-02",
			},
		}

		h.dataAccessor.EXPECT().getUserList().Return(users, nil)

		actual, actualErr := h.usecase.getUserList()
		if actualErr != nil {
			t.Fatal(actualErr)
		}
		if !reflect.DeepEqual(actual, users) {
			t.Fatal(actual, actualErr)
		}
	})
	t.Run("異常系", func(t *testing.T) {
		t.Parallel()
		h := newUsecaseTestHelper(t)
		defer h.ctrl.Finish()
		err := errors.New("error")

		h.dataAccessor.EXPECT().getUserList().Return([]getUserListResult{}, err)

		if _, actualErr := h.usecase.getUserList(); actualErr != err {
			t.Fatal(actualErr)
		}
	})

}
