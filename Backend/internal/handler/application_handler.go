package handler

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/hanoguz056-hub/jobboard/internal/service"
)

type ApplicationHandler struct {
	service *service.ApplicationService
}

func NewApplicationHandler(service *service.ApplicationService) (*ApplicationHandler) {
	return &ApplicationHandler{service: service}
}

// Apply godoc
// @Summary      Откликнуться на вакансию
// @Tags         applications
// @Security     BearerAuth
// @Accept       mpfd
// @Produce      json
// @Param        id            path      string  true   "ID вакансии"
// @Param        resume        formData  file    true   "PDF резюме"
// @Param        cover_letter  formData  string  false  "Сопроводительное письмо"
// @Success      201           {object}  models.Application
// @Failure      400           {object}  map[string]string
// @Router       /api/v1/jobs/{id}/apply [post]
func (h *ApplicationHandler) Apply(c *fiber.Ctx) error {
	file, err := c.FormFile("resume")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "файл не найден"})
	}

	src, err := file.Open()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Не удалось открыть файл"})
	}

	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	
	if err != nil {
    return c.Status(400).JSON(fiber.Map{"error": "не удалось прочитать файл"})

	}

	coverLetter := c.FormValue("cover_letter")
	jobID := c.Params("id")
	candidateID := c.Locals("user_id").(string)

	result, err := h.service.Apply(c.Context(), jobID, candidateID, coverLetter, fileBytes, file.Filename)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(result)
}

// GetMyApplications godoc
// @Summary      Мои отклики
// @Tags         applications
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}   models.Application
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/applications/my [get]
func(h *ApplicationHandler) GetMyApplications(c *fiber.Ctx) error {
	candidateID := c.Locals("user_id").(string)

	result, err := h.service.GetMyApplications(c.Context(), candidateID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(result)
}

// GetJobApplications godoc
// @Summary      Отклики на вакансию
// @Tags         applications
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "ID вакансии"
// @Success      200  {array}   models.Application
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/jobs/{id}/applications [get]
func(h *ApplicationHandler) GetJobApplications(c *fiber.Ctx) error {
	jobID := c.Params("id")
	result, err := h.service.GetJobApplications(c.Context(), jobID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(result)
}

// UpdateApplicationStatus godoc
// @Summary      Изменить статус отклика
// @Tags         applications
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "ID отклика"
// @Param        input  body      object  true  "Статус"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/applications/{id}/status [patch]
func(h *ApplicationHandler) UpdateApplicationStatus(c *fiber.Ctx) error {
	var payload struct {
		Status string `json:"status"`
	}
	id := c.Params("id")

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	 err := h.service.UpdateApplicationStatus(c.Context(), id, payload.Status)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "updated"})
}