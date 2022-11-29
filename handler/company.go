package handler

import (
	"companies-crud/model"
	"companies-crud/service"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Company(s service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		companyId := c.Param("companyId")

		company, err := s.Company(companyId)
		if errors.Is(err, model.ErrCompanyNotFound) {
			return c.JSON(http.StatusNotFound, model.ErrCompanyNotFound.Error())
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get company")
		}

		return c.JSON(http.StatusOK, company)
	}
}

func CreateCompany(s service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		company := new(model.CompanyDto)
		if err := c.Bind(&company); err != nil {
			return c.JSON(http.StatusBadRequest, "unable to bind json")
		}

		err := s.CreateCompany(*company)
		if errors.Is(err, model.ErrCompanyAlreadyExists) {
			return c.JSON(http.StatusForbidden, model.ErrCompanyAlreadyExists.Error())
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to insert company")
		}

		return c.JSON(http.StatusCreated, echo.Map{"created": "1"})

	}
}

func PatchCompany(s service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		companyId := c.Param("companyId")

		company, err := s.Company(companyId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to get company")
		}

		err = json.NewDecoder(c.Request().Body).Decode(&company)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "unable to bind json")
		}

		err = s.UpdateCompany(company)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "something unexpected happened")
		}

		return c.JSON(http.StatusOK, echo.Map{"updated": "1"})
	}
}

func DeleteCompany(s service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		companyId := c.Param("companyId")

		err := s.DeleteCompany(companyId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "something unexpected happened")
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
