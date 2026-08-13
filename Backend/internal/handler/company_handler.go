package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/service"
)

type CompanyHandler struct {
	service *service.CompanyService
}

func NewCompanyHandler(service *service.CompanyService) (*CompanyHandler) {
	return &CompanyHandler{
		service: service,
	}
}


// GetCompanies godoc
// @Summary      Список компаний
// @Tags         companies
// @Produce      json
// @Param        page   query     int  false  "Страница"
// @Param        limit  query     int  false  "Лимит"
// @Success      200    {array}   models.Company
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/companies [get]
func(h *CompanyHandler) GetCompanies(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)

	limit := c.QueryInt("limit", 10)

	result, err := h.service.GetCompanies(c.Context(), page, limit)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(result)
}


// GetCompanyByID godoc
// @Summary      Компания по ID
// @Tags         companies
// @Produce      json
// @Param        id   path      string  true  "ID компании"
// @Success      200  {object}  models.Company
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/companies/{id} [get]
func(h *CompanyHandler) GetCompanyByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.service.GetCompanyByID(c.Context(), id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}


// CreateCompany godoc
// @Summary      Создать компанию
// @Tags         companies
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        input  body      models.Company  true  "Данные компании"
// @Success      200    {object}  models.Company
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/companies [post]
func(h *CompanyHandler) CreateCompany(c *fiber.Ctx) error {
	var company models.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	ownerID := c.Locals("user_id").(string)

	result, err := h.service.CreateCompany(c.Context(), company, ownerID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}


// UpdateCompany godoc
// @Summary      Обновить компанию
// @Tags         companies
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id     path      string          true  "ID компании"
// @Param        input  body      models.Company  true  "Данные компании"
// @Success      200    {object}  models.Company
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/companies/{id} [put]
func(h *CompanyHandler) UpdateCompany(c *fiber.Ctx) error {
	var company models.Company
	
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	company.ID = id

	ownerID := c.Locals("user_id").(string)

	err := h.service.UpdateCompany(c.Context(), company, ownerID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(company)
}