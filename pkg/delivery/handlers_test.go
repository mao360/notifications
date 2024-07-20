package delivery

import (
	"context"
	"errors"
	"github.com/mao360/notifications/models"
	"github.com/mao360/notifications/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Registration(t *testing.T) {

	TestCases := []struct {
		testName     string
		userStruct   models.User
		inputBody    string
		expectedBody string
		expectedCode int
	}{
		{
			testName: "ok",
			userStruct: models.User{
				UserName:    "test",
				Password:    "qwerty",
				DateOfBirth: "1970-01-01",
			},
			inputBody:    `{"user_name":"test","password":"qwerty","date_of_birth":"1970-01-01"}`,
			expectedBody: ``,
			expectedCode: 201,
		},
		{
			testName: "wrong body",
			userStruct: models.User{
				UserName:    "test",
				Password:    "qwerty",
				DateOfBirth: "1970-01-01",
			},
			inputBody:    `{"user_name:"test","password":"qwerty","date_of_birth":"1970-01-01"}`,
			expectedBody: `{"message":"unmarshal error"}`,
			expectedCode: 500,
		},
		{
			testName: "wrong user",
			userStruct: models.User{
				Password:    "qwerty",
				DateOfBirth: "1970-01-01",
			},
			inputBody:    `{"password":"qwerty","date_of_birth":"1970-01-01"}`,
			expectedBody: `{"message":"newUser error"}`,
			expectedCode: 500,
		},
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	for _, v := range TestCases {
		t.Run(v.testName, func(t *testing.T) {
			service := mocks.NewServiceI(t)
			service.On("NewUser", mock.Anything, &v.userStruct).
				Maybe().
				Return(func(context.Context, *models.User) error {
					var err error
					switch {
					case v.testName == "wrong user":
						err = errors.New("not enough data for create user")
					case v.testName == "wrong body":
						err = errors.New("unmarshal error")
					}
					return err
				})

			h := &Handler{
				service: service,
				sugared: sugar,
			}

			handler := http.HandlerFunc(h.Registration)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/reg", strings.NewReader(v.inputBody))
			handler.ServeHTTP(w, r)
			assert.Equal(t, v.expectedBody, w.Body.String())
			assert.Equal(t, v.expectedCode, w.Code)
		})
	}
}

func TestHandler_Authorization(t *testing.T) {

	TestCases := []struct {
		testName     string
		userName     string
		password     string
		inputBody    string
		expectedBody string
		expectedCode int
	}{
		{
			testName:     "ok",
			userName:     "test",
			password:     "qwerty",
			inputBody:    `{"user_name":"test","password":"qwerty"}`,
			expectedBody: `{"token":"someJWT"}`,
			expectedCode: 200,
		},
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	for _, v := range TestCases {
		t.Run(v.testName, func(t *testing.T) {
			service := mocks.NewServiceI(t)
			service.On("GetUser", context.Background(), v.userName, v.password).
				Once().
				Return(&models.User{
					UserName:    v.userName,
					Password:    v.password,
					DateOfBirth: "10-10-1990",
				}, nil)
			service.On("GenerateToken", context.Background(), v.userName, v.password).
				Once().
				Return("someJWT", nil)

			h := &Handler{
				service: service,
				sugared: sugar,
			}

			handler := http.HandlerFunc(h.Authorization)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/auth", strings.NewReader(v.inputBody))
			handler.ServeHTTP(w, r)
			assert.Equal(t, v.expectedBody, w.Body.String())
			assert.Equal(t, v.expectedCode, w.Code)
		})
	}
}
