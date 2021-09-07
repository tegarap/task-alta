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
	if err = config.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllBookController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataBookForGetBooks()

	type BookResponse struct {
		Status string        `json:"status"`
		Data   []models.Book `json:"data"`
	}

	book := models.Book{
		Title:     "Buku Pengantar Tidur",
		Author:    "Prabu Tegar",
		Publisher: "Maliki Press",
	}


	for i := 1; i <= 2; i++ {
		if i == 2 {
			//drop table
			config.DB.Migrator().DropTable(&models.Book{})
		}
		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		GetAllBookController(context)


		var response BookResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if i == 1 {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, book.Title, response.Data[0].Title)
			assert.Equal(t, book.Author, response.Data[0].Author)
			assert.Equal(t, book.Publisher, response.Data[0].Publisher)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestGetBookController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataBookForGetBooks()
	idParam := []string{"1", "1a", "99"}

	type BookResponse struct {
		Status string      `json:"status"`
		Data   models.Book `json:"data"`
	}

	book := models.Book{
		Title:     "Buku Pengantar Tidur",
		Author:    "Prabu Tegar",
		Publisher: "Maliki Press",
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		GetBookController(context)

		var response BookResponse
		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if value == "1" {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, book.Title, response.Data.Title)
			assert.Equal(t, book.Author, response.Data.Author)
			assert.Equal(t, book.Publisher, response.Data.Publisher)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestCreateBookController(t *testing.T) {
	e := initEchoTestAPI()

	type BookResponse struct {
		Status string      `json:"status"`
		Data   models.Book `json:"data"`
	}

	var response BookResponse

	for i := 1; i <= 2; i++ {
		reqBody := make(map[string]interface{})
		if i == 1 {
			reqBody["title"] = "Madangi Sing Peteng"
			reqBody["author"] = "Songgeno"
			reqBody["publisher"] = "Luwak Publishing"
		} else {
			reqBody["title"] = ""
		}

		requestBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		CreateBookController(context)

		resBody := res.Body.String()

		json.Unmarshal([]byte(resBody), &response)

		if i == 1 {
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "success", response.Status)
			assert.Equal(t, "Madangi Sing Peteng", response.Data.Title)
			assert.Equal(t, "Songgeno", response.Data.Author)
			assert.Equal(t, "Luwak Publishing", response.Data.Publisher)
		} else {
			assert.Equal(t, http.StatusBadRequest, res.Code)
			assert.Equal(t, "fail", response.Status)
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataBookForGetBooks()
	idParam := []string{"1", "1a", "99"}

	type BookResponse struct {
		Status string      `json:"status"`
		Data   models.Book `json:"data"`
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		DeleteBookController(context)

		var response BookResponse
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

func TestUpdateBookController(t *testing.T) {
	e := initEchoTestAPI()
	InsertDataBookForGetBooks()
	idParam := []string{"1", "1a", "99"}

	requestBody, _ := json.Marshal(map[string]string{
		"author":  "Ini Author",
		"title": "Title",
	})

	type BookResponse struct {
		Status string      `json:"status"`
		Data   models.Book `json:"data"`
	}

	for _, value := range idParam {
		req := httptest.NewRequest(http.MethodPut, "/books/:id", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetParamNames("id")
		context.SetParamValues(value)
		UpdateBookController(context)

		var response BookResponse
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
