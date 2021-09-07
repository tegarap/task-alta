package user

import (
	"book-api-mvc/config"
	"book-api-mvc/models"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func initEchoTestAPI() *echo.Echo {
	config.InitDBTest("user")
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := models.User{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllUserController(t *testing.T) {
	var testCases = []struct {
		expectCode   int
		responStatus string
	}{
		{
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	for i, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if i == 1 {
			//drop table
			config.DB.Migrator().DropTable(&models.User{})
		}
		if assert.NoError(t, GetAllUserController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		idParam      string
		expectCode   int
		responStatus string
	}{
		{
			idParam:      "1",
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			idParam:      "1a",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			idParam:      "99",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "99" {
			//drop table
			config.DB.Migrator().DropTable(&models.User{})
		}
		if assert.NoError(t, GetUserController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}


func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		reqBody      map[string]string
		expectCode   int
		responStatus string
	}{
		{
			reqBody:      map[string]string{"name": "Test Nama", "email": "iniemail@test.com", "password": "luwakwhitecoffe"},
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			reqBody:      map[string]string{"name": ""},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, CreateUserController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []struct {
		idParam      string
		expectCode   int
		responStatus string
	}{
		{
			idParam:      "1",
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			idParam:      "1a",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			idParam:      "99",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			idParam:      "drop",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "drop" {
			//drop table
			config.DB.Migrator().DropTable(&models.User{})
		}
		if assert.NoError(t, DeleteUserController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		idParam      string
		expectCode   int
		responStatus string
	}{
		{
			idParam:      "1",
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			idParam:      "1a",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			idParam:      "99",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			idParam:      "drop",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	requestBody, _ := json.Marshal(map[string]string{
		"name":  "Tes Nama",
		"email": "iniemail@test.com",
	})
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "drop" {
			//drop table
			config.DB.Migrator().DropTable(&models.User{})
		}
		if assert.NoError(t, UpdateUserController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestLoginUsersController(t *testing.T) {
	var testCases = []struct {
		reqBody      map[string]string
		expectCode   int
		responStatus string
	}{
		{
			reqBody:      map[string]string{"email": "alta@gmail.com", "password": "123"},
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			reqBody:      map[string]string{"email": "kliru@gmail.com", "password": "123123123"},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertDataUserForGetUsers()
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, LoginUsersController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}
