package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/service"
)



type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(service *service.JobService) (*JobHandler) {
	return &JobHandler{
		service: service,
	}
}


// GetJobs godoc
// @Summary      Список вакансий
// @Tags         jobs
// @Produce      json
// @Param        page        query     int     false  "Страница"
// @Param        limit       query     int     false  "Лимит"
// @Param        search      query     string  false  "Поиск по названию"
// @Param        city        query     string  false  "Город"
// @Param        type        query     string  false  "Тип: full/part/remote"
// @Param        salary_min  query     number  false  "Мин зарплата"
// @Param        salary_max  query     number  false  "Макс зарплата"
// @Success      200         {array}   models.Job
// @Failure      400         {object}  map[string]string
// @Router       /api/v1/jobs [get]
func(h *JobHandler) GetJobs(c *fiber.Ctx) error {
	var filter models.JobFilter
	
	page := c.QueryInt("page", 1)

	limit := c.QueryInt("limit", 10)

	filter.Page = page
	filter.Limit = limit

	result, err := h.service.GetJobs(c.Context(), filter)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(result)
}


// GetJobByID godoc
// @Summary      Вакансия по ID
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "ID вакансии"
// @Success      200  {object}  models.Job
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/jobs/{id} [get]
func(h *JobHandler) GetJobByID(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.service.GetJobByID(c.Context(), id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(result)
}

// CreateJob godoc
// @Summary      Создать вакансию
// @Tags         jobs
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        input  body      models.Job  true  "Данные вакансии"
// @Success      201    {object}  models.Job
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/jobs [post]
func(h *JobHandler) CreateJob(c *fiber.Ctx) error {
	var job models.Job

	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	companyID := job.CompanyID
	
	result, err := h.service.CreateJob(c.Context(), job, companyID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(201).JSON(result)
}

// UpdateJob godoc
// @Summary      Обновить вакансию
// @Tags         jobs
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id     path      string      true  "ID вакансии"
// @Param        input  body      models.Job  true  "Данные вакансии"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/jobs/{id} [put]
func(h *JobHandler) UpdateJob(c *fiber.Ctx) error {
	
	var job models.Job

	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	id := c.Params("id")
	userID := c.Locals("user_id").(string)
	

	job.ID = id
	
	 err := h.service.UpdateJob(c.Context(), job, userID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Updated"})

}

// UpdateJobStatus godoc
// @Summary      Изменить статус вакансии
// @Tags         jobs
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "ID вакансии"
// @Param        input  body      object  true  "Статус"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/jobs/{id}/status [patch]
func(h *JobHandler) UpdateJobStatus(c *fiber.Ctx) error {
	var payload struct {
		Status string `json:"status"`
	}


	id := c.Params("id")
	userID := c.Locals("user_id").(string)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}
	

	err := h.service.UpdateJobStatus(c.Context(), id, payload.Status, userID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "statu updated"})
}

// DeleteJob godoc
// @Summary      Удалить вакансию
// @Tags         jobs
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "ID вакансии"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/jobs/{id} [delete]
func(h *JobHandler) DeleteJob(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := c.Locals("user_id").(string)
	role := c.Locals("role").(string)

	err := h.service.DeleteJob(c.Context(), id, userID, role)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "deleted"})
}

// ApproveJob godoc
// @Summary      Одобрить вакансию (Admin)
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "ID вакансии"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/admin/jobs/{id}/approve [patch]
func(h *JobHandler) ApproveJob(c *fiber.Ctx) error {
	jobID := c.Params("jobID")

	err := h.service.ApproveJob(c.Context(), jobID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "approved"})
}