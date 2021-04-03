package usercommand

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
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

func TestServerCreate(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()
		h := newServerTestHelper(t)

		f := make(url.Values)
		f.Set("email", "test@gmail.com")
		f.Set("password", "password")
		f.Set("gender", "男")

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

		rec := httptest.NewRecorder()
		ctx := h.etx.NewContext(req, rec)

		h.usecase.EXPECT().Create(createRequest{
			Email:    "test@gmail.com",
			Password: "password",
			Gender:   "男"}).Return(nil)

		if assert.NoError(t, h.server.Create(ctx)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
}
