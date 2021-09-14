package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"verifyMyTest/m/app/http/controllers"
	"verifyMyTest/m/app/models"
	"verifyMyTest/m/routes"
	"verifyMyTest/m/util"
)

type RequestTestParameters struct {
	Url     string
	Method  string
	Message string
	Success bool
	Status  int
}

func setupRouter() *gin.Engine {
	r := routes.Setup()
	return r
}

var routeCases = []struct {
	name       string
	parameters *RequestTestParameters
	body       io.Reader
}{
	{"get customers", &RequestTestParameters{
		Url:     "/customers",
		Method:  "GET",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusOK,
	}, nil},
	{"get customer record found", &RequestTestParameters{
		Url:     "/customer/id/1",
		Method:  "GET",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusOK,
	}, nil},
	{"get customer no arg", &RequestTestParameters{
		Url:     "/customer/asdasd",
		Method:  "GET",
		Message: "Not found",
		Success: false,
		Status:  http.StatusNotFound,
	}, nil},
	{"get customer no records found", &RequestTestParameters{
		Url:     "/customer/id/1111111",
		Method:  "GET",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusOK,
	}, nil},
	{"get customer records found", &RequestTestParameters{
		Url:     "/customer/id/1",
		Method:  "GET",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusOK,
	}, nil},
	{"create customer", &RequestTestParameters{
		Url:     "/customer",
		Method:  "POST",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusCreated,
	}, marshalCreateCustomer()},
	{"delete customer not found", &RequestTestParameters{
		Url:     "/customer/123131",
		Method:  "DELETE",
		Message: controllers.RecordNotFound,
		Success: false,
		Status:  http.StatusBadRequest,
	}, nil},
}

func TestRestRoutes(t *testing.T) {
	for _, testCase := range routeCases {
		t.Run(testCase.name, func(t *testing.T) {
			assertReturn(t, testCase.parameters, testCase.body)
		})
	}
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestNonEczisteRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/issononecziste", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)

}

func marshalCreateCustomer() io.Reader {
	str := strconv.Itoa(int(rand.Float64()))
	s := fmt.Sprintf("email%s@test.com", str)

	newCustomer := models.Customer{
		Name:      "MICHAEL J FOX",
		Password:  s,
		Age:       12,
		Email:     s,
		Address:   "asdasd",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	customerJSON, _ := json.Marshal(newCustomer)

	return bytes.NewBuffer(customerJSON)

}
func TestShouldCreateCustomer(t *testing.T) {

	str := strconv.Itoa(int(rand.Float64()))
	s := fmt.Sprintf("email%s@test.com", str)

	newCustomer := models.Customer{
		Name:      "MICHAEL J FOX JR",
		Password:  s,
		Age:       12,
		Email:     s,
		Address:   "asdasd",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	customerJSON, err := json.Marshal(newCustomer)

	if err != nil {
		t.Error(err.Error())
	}

	parameters := RequestTestParameters{
		Url:     "/customer",
		Method:  "POST",
		Message: controllers.OkMessage,
		Success: true,
		Status:  http.StatusCreated,
	}
	assertReturn(t, &parameters, bytes.NewBuffer(customerJSON))

}

func assertReturn(t *testing.T, parameters *RequestTestParameters, body io.Reader) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(parameters.Method, parameters.Url, body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	data := controllers.Response{Data: util.Data{}}
	err := json.Unmarshal([]byte(w.Body.String()), &data)
	assert.Nil(t, err)
	assert.Equal(t, parameters.Message, data.Message)
	assert.Equal(t, parameters.Success, data.Success)
	assert.Equal(t, parameters.Status, w.Code)
}
