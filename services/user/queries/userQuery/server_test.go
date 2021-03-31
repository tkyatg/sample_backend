package userquery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type (
	serverTestHelper struct {
		t       *testing.T
		ctrl    *gomock.Controller
		etx     *echo.Echo
		server  *server
		usecase *MockUsecase
	}
)

func newServerTestHelper(t *testing.T) *serverTestHelper {
	ctrl := gomock.NewController(t)
	usecase := NewMockUsecase(ctrl)
	etx := echo.New()
	server := &server{usecase}
	return &serverTestHelper{t, ctrl, etx, server, usecase}
}

func (s *serverTestHelper) asExpected(expect interface{}, actual string) (bool, error) {
	var actualJson bytes.Buffer
	json.Compact(&actualJson, []byte(actual))

	expByte, err := json.Marshal(expect)
	if err != nil {
		return false, err
	}
	var expectJson bytes.Buffer
	json.Compact(&expectJson, []byte(expByte))
	return assert.Equal(s.t, expectJson, actualJson), nil
}

func TestServerGetUserList(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newServerTestHelper(t)
		expect := []getUserListResult{
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
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		ctx := h.etx.NewContext(req, rec)
		ctx.SetPath("/users")

		h.usecase.EXPECT().getUserList().Return(expect, nil)

		if assert.NoError(t, h.server.GetUserList(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			h.asExpected(expect, rec.Body.String())
		}
	})
}

func TestServerGetUserByID(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newServerTestHelper(t)

		userUUID := uuid.New().String()

		expect := getUserByIDResult{
			UserUUID:         "UserUUID-01",
			DisplayName:      "DisplayName-01",
			BirthDay:         "BirthDay-01",
			Gender:           "Gender-01",
			ImageURL:         "ImageURL-01",
			FreeTime:         "FreeTime-01",
			SelfIntroduction: "SelfIntroduction-01",
			CreatedAt:        "CreatedAt-01",
			UpdatedAt:        "UpdatedAt-01",
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		ctx := h.etx.NewContext(req, rec)
		ctx.SetPath("/users/:uuid")
		ctx.SetParamNames("uuid")
		ctx.SetParamValues(userUUID)

		h.usecase.EXPECT().getUserByID(userUUID).Return(expect, nil)

		if assert.NoError(t, h.server.GetUserByID(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			h.asExpected(expect, rec.Body.String())
		}
	})
}
