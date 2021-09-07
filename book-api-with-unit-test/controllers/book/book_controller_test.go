package book

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

type bookResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func initEchoTestAPI() *echo.Echo {
	config.InitDBTest("book")
	e := echo.New()
	return e
}

func InsertDataBookForGetBooks() error {
	book := models.Book{
		Title:     "Buku Pengantar Tidur",
		Author:    "Prabu Tegar",
		Publisher: "Maliki Press",
	}

	var err error
	if err = config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllBookController(t *testing.T) {
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
	InsertDataBookForGetBooks()

	for i, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if i == 1 {
			//drop table
			config.DB.Migrator().DropTable(&models.Book{})
		}
		if assert.NoError(t, GetAllBookController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response bookResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestGetBookController(t *testing.T) {
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
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "99" {
			//drop table
			config.DB.Migrator().DropTable(&models.Book{})
		}
		if assert.NoError(t, GetBookController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response bookResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestCreateBookController(t *testing.T) {
	var testCases = []struct {
		reqBody      map[string]string
		expectCode   int
		responStatus string
	}{
		{
			reqBody:      map[string]string{"title":"Madangi Sing Peteng", "author":"Songgeno", "publisher":"Luwak Publishing"},
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			reqBody:      map[string]string{"title": ""},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, CreateBookController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response bookResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestDeleteBookController(t *testing.T) {
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
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "drop" {
			//drop table
			config.DB.Migrator().DropTable(&models.Book{})
		}
		if assert.NoError(t, DeleteBookController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response bookResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestUpdateBookController(t *testing.T) {
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
	InsertDataBookForGetBooks()

	requestBody, _ := json.Marshal(map[string]string{
		"author": "Ini Author",
		"title":  "Title",
	})
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/books/:id", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(testCase.idParam)
		if testCase.idParam == "drop" {
			//drop table
			config.DB.Migrator().DropTable(&models.Book{})
		}
		if assert.NoError(t, UpdateBookController(context)) {
			assert.Equal(t, testCase.expectCode, res.Code)
			resBody := res.Body.String()

			var response bookResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}
