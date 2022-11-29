package route

import (
	"companies-crud/handler"
	"companies-crud/service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, s service.Service) {
	g := e.Group("/api/company")

	g.GET("/:companyId", handler.Company(s))
	g.PATCH("/:companyId", handler.PatchCompany(s))
	g.DELETE("/:companyId", handler.DeleteCompany(s))
	g.POST("", handler.CreateCompany(s))

}
