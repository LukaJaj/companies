package handler_test

import (
	"companies-crud/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompany(t *testing.T) {
	getCompany := handler.Company(*mockService)

	t.Run("GET Company", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/company", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("0104c7e9-03b6-4d2f-a725-03730244d8c9")

		testFile, _ := filepath.Abs("../mock/json/company.json")

		body, _ := ioutil.ReadFile(testFile)

		if assert.NoError(t, getCompany(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, FormatExpected(body), FormatActual(rec.Body.String()))
		}
	})

	t.Run("Error GET Company", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/company", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("error_company")

		if assert.NoError(t, getCompany(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "\"failed to get company\"\n", rec.Body.String())
		}
	})

	t.Run("Company Not Found", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/company", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("company_not_found")

		if assert.NoError(t, getCompany(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "\"company with given id not found\"\n", rec.Body.String())
		}
	})
}

func TestCreateCompany(t *testing.T) {
	createCompany := handler.CreateCompany(*mockService)

	t.Run("GET Company", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/company.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodGet, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, createCompany(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "{\"created\":\"1\"}", FormatActual(rec.Body.String()))
		}
	})

	t.Run("Unable To Bind Company", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/badRequestCompany.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodGet, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, createCompany(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "\"unable to bind json\"\n", rec.Body.String())
		}
	})

	t.Run("Company Already Exists", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/newCompany.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodGet, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, createCompany(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
			assert.Equal(t, "\"company with given id already exists\"\n", rec.Body.String())
		}
	})

	t.Run("Company Insertion DB Error", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/insertionDbErrorCompany.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodGet, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, createCompany(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "\"failed to insert company\"\n", rec.Body.String())
		}
	})
}

func TestUpdateCompany(t *testing.T) {
	updateCompany := handler.PatchCompany(*mockService)

	t.Run("Update Company Error While GET", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/company.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodGet, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("5504c7e9-45yu-4d2f-a725-03730244dvc8")

		if assert.NoError(t, updateCompany(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "\"failed to get company\"\n", rec.Body.String())
		}
	})

	t.Run("Update Company Error", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/updateCompanyError.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodPatch, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("9004c7e9-03b6-4d2f-a725-03730244d8d7")

		if assert.NoError(t, updateCompany(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "\"something unexpected happened\"\n", rec.Body.String())
		}
	})

	t.Run("Update Company", func(t *testing.T) {
		testFile, _ := filepath.Abs("../mock/json/updateCompany.json")
		body, _ := ioutil.ReadFile(testFile)

		req := httptest.NewRequest(http.MethodPatch, "/api/company", strings.NewReader(string(body)))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("0204c7e9-03b6-4d2f-a725-03730244d8c9")

		if assert.NoError(t, updateCompany(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "{\"updated\":\"1\"}\n", rec.Body.String())
		}
	})

}
func TestRemoveCompany(t *testing.T) {

	deleteCompany := handler.DeleteCompany(*mockService)

	t.Run("Delete Company", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/api/company", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("2704c7e9-03b6-4d2f-a725-03730244d8k1")

		if assert.NoError(t, deleteCompany(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
			assert.Equal(t, "null\n", rec.Body.String())
		}
	})

	t.Run("DB Error While Deleting Company", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodDelete, "/api/company", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:companyId")
		c.SetParamNames("companyId")
		c.SetParamValues("error")

		if assert.NoError(t, deleteCompany(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "\"something unexpected happened\"\n", rec.Body.String())
		}
	})

}
