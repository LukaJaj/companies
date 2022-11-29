package handler_test

import (
	"bytes"
	"companies-crud/mock"
	"companies-crud/service"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var mockService *service.Service
var e *echo.Echo

func TestMain(m *testing.M) {
	e = echo.New()
	db := gorm.DB{}
	mockStore := mock.New(&db)
	mockService = service.New(mockStore)

	m.Run()
	os.Exit(0)
}

func FormatActual(jsonString string) string {
	formatted := strings.TrimSuffix(jsonString, "\n")
	return formatted
}

func FormatExpected(data []byte) string {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, data); err != nil {
		panic(err)
	}
	return buffer.String()
}
