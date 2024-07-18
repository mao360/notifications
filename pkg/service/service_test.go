package service

import (
	"github.com/mao360/notifications/models"
	"github.com/mao360/notifications/pkg/repo/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GenerateToken(t *testing.T) {

	TestCases := []struct {
		testName string
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

	for _, v := range TestCases {
		t.Run(v.testName, func(t *testing.T) {
			repository := mocks.NewRepoI(t)
			repository.On("GetUser").
				Once().
				Return(models.User{})
			assert.Equal(t, v.expectedBody, w.Body.String())
			assert.Equal(t, v.expectedCode, w.Code)
		})
	}
}

func TestService_ParseToken(t *testing.T) {

	TestCases := []struct {
		testName string
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

	for _, v := range TestCases {
		t.Run(v.testName, func(t *testing.T) {
			assert.Equal(t, v.expectedBody, w.Body.String())
			assert.Equal(t, v.expectedCode, w.Code)
		})
	}
}
