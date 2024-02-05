package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test/entity"
	mock "test/mock/mockUsecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UserSignup(t *testing.T) {
	testCase := map[string]struct {
		input         entity.CreateUserInput
		buildStub     func(useCaseMock *mock.MockRegisterUserUseCase, signupData entity.CreateUserInput)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"Valid Signup": {
			input: entity.CreateUserInput{
				FirstName: "akhil",
				LastName:  "c",
				Email:     "akhilc89@gmail.com",
				Gender:    "MALE",
			},
			buildStub: func(useCaseMock *mock.MockRegisterUserUseCase, signupData entity.CreateUserInput) {
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}
				useCaseMock.EXPECT().Execute(gomock.Any(), signupData).Times(1).Return(entity.User{
					FirstName: "akhil",
					LastName:  "c",
					Email:     "akhilc89@gmail.com",
					Gender:    "MALE",
				}, nil)

			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, responseRecorder.Code)
			},
		},
		"user couldnot sign up": {
			input: entity.CreateUserInput{
				FirstName: "akhil",
				LastName:  "c",
				Email:     "akhilc89@gmail.com",
				Gender:    "MALE",
			},
			buildStub: func(useCaseMock *mock.MockRegisterUserUseCase, signupData entity.CreateUserInput) {
				// copying signupData to domain.user for pass to Mock usecase
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}
				useCaseMock.EXPECT().Execute(gomock.Any(), signupData).Times(1).Return(entity.User{}, errors.New("cannot sign up"))
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}
	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mock.NewMockRegisterUserUseCase(ctrl)
			test.buildStub(mockUseCase, test.input)
			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", userHandler.Register)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()
			server.ServeHTTP(responseRecorder, mockRequest)

			test.checkResponse(t, responseRecorder)

		})
	}
}
