package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hanoguz056-hub/jobboard/internal/models"
	"github.com/hanoguz056-hub/jobboard/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler (service *service.UserService) (*UserHandler) {
	return &UserHandler{service: service}
}


type BanUserRequest struct {
	IsBanned bool `json:"is_banned"`
}

// Register godoc
// @Summary      Регистрация нового пользователя
// @Description  Создает новый аккаунт пользователя с указанной ролью (candidate, employer)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      models.User  true  "Данные для регистрации"
// @Success      201    {object}  models.User        "Пользователь успешно создан"
// @Failure      400    {object}  map[string]string  "Неверный формат запроса или валидации"
// @Failure      409    {object}  map[string]string  "Пользователь с таким email уже существует"
// @Failure      500    {object}  map[string]string  "Внутренняя ошибка сервера"
// @Router       /api/v1/auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.Register(c.Context(), req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(result)
}



// Login godoc
// @Summary      Вход в систему
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      models.LoginRequest  true  "Данные пользователя"
// @Success      200    {object}  models.AuthResponse
// @Failure      401    {object}  map[string]string
// @Router       /api/v1/auth/login [post]
func(h *UserHandler) Login(c *fiber.Ctx) error {
	var input models.LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	resp, err := h.service.Login(c.UserContext(), input )	

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(resp)
}


// GetAllUsers godoc
// @Summary      Список всех пользователей
// @Tags         admin
// @Security     BearerAuth
// @Produce      json
// @Param        page   query     int  false  "Страница"
// @Param        limit  query     int  false  "Лимит"
// @Success      200    {array}   models.User
// @Failure      500    {object}  map[string]string
// @Router       /api/v1/admin/users [get]
func(h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)

	limit := c.QueryInt("limit", 10)

	result, err := h.service.GetAllUsers(c.Context(), page, limit)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error"})
	}

	return c.Status(200).JSON(result)
}


// BanUser godoc
// @Summary      Заблокировать пользователя
// @Tags         admin
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id     path      string          true  "ID пользователя"
// @Param        input  body      BanUserRequest  true  "Данные"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Router       /api/v1/admin/users/{id}/ban [patch]
func(h *UserHandler) BanUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	
	var req BanUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	err := h.service.BanUser(c.Context(), userID, req.IsBanned)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "updated"})
}