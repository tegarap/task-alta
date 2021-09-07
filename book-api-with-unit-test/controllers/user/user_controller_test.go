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
	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	type UserResponse struct {
		Status string        `json:"status"`
		Data   []models.User `json:"data"`
	}

	user := models.User{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	for i := 1; i <= 2; i++ {
		if i == 2 {
			//drop table
			config.DB.Migrator().DropTable(&models.User{})
		}
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		GetAllUserController(context)


		var response UserResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if i == 1 {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, user.Name, response.Data[0].Name)
			assert.Equal(t, user.Email, response.Data[0].Email)
			assert.Equal(t, user.Password, response.Data[0].Password)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestGetUserController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataUserForGetUsers()
	idParam := []string{"1", "1a", "99"}

	type UserResponse struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}

	user := models.User{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		GetUserController(context)

		var response UserResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if value == "1" {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, user.Name, response.Data.Name)
			assert.Equal(t, user.Email, response.Data.Email)
			assert.Equal(t, user.Password, response.Data.Password)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestCreateUserController(t *testing.T) {
	e := initEchoTestAPI()

	type UserResponse struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}

	var response UserResponse

	for i := 1; i <= 2; i++ {
		reqBody := make(map[string]interface{})
		if i == 1 {
			reqBody["name"] = "Test Nama"
			reqBody["email"] = "iniemail@test.com"
			reqBody["password"] = "luwakwhitecoffe"
		} else {
			reqBody["name"] = ""
		}

		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		CreateUserController(context)

		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if i == 1 {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "Test Nama", response.Data.Name)
			assert.Equal(t, "iniemail@test.com", response.Data.Email)
			assert.Equal(t, "luwakwhitecoffe", response.Data.Password)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataUserForGetUsers()
	idParam := []string{"1", "1a", "99"}

	type UserResponse struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		DeleteUserController(context)

		var response UserResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if value == "1" {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataUserForGetUsers()
	idParam := []string{"1", "1a", "99"}

	requestBody, _ := json.Marshal(map[string]string{
		"name":  "Tes Nama",
		"email": "iniemail@test.com",
		//"password": "luwakwhitecoffe",
	})

	type UserResponse struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		UpdateUserController(context)

		var response UserResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if value == "1" {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestLoginUsersController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataUserForGetUsers()

	type UserResponse struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}

	for i := 1; i <= 2; i++ {
		reqBody := make(map[string]string)
		if i == 1 {
			reqBody["email"] = "alta@gmail.com"
			reqBody["password"] = "123"
		} else {
			reqBody["email"] = "kliru@gmail.com"
			reqBody["password"] = "1234123"
		}

		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		LoginUsersController(context)

		var response UserResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if i == 1 {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}
