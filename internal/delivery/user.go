package delivery

import (
	"fmt"
	"rbac/domain"
	usecase "rbac/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.UserUsecase.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: users,
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	params := c.AllParams()
	name, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: "get pararm failed",
		})
	}
	name, err := h.UserUsecase.GetUser(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	if name == "" {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("user %s not found", name),
		})
	}
	return nil
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: "get pararm failed",
		})
	}
	err := h.UserUsecase.DeleteUser(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}

func (h *UserHandler) AddRole(c *fiber.Ctx) error {
	type request struct {
		UserName string `json:"user_name"`
		RoleName string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	// Call the usecase method to add user to role
	err := h.UserUsecase.AddRole(req.UserName, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *UserHandler) RemoveRole(c *fiber.Ctx) error {
	type request struct {
		UserName string `json:"user_name"`
		RoleName string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	// Call the usecase method to remove user from role
	err := h.UserUsecase.RemoveRole(req.UserName, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *UserHandler) FindAllObjectRelations(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	relations, err := h.UserUsecase.FindAllObjectRelations(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: relations,
	})
}

func (h *UserHandler) AddRelation(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		UserName        string `json:"user_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	// Call the usecase method to add permission to user
	err := h.UserUsecase.AddRelation(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *UserHandler) RemoveRelation(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		UserName        string `json:"user_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	// Call the usecase method to add permission to user
	err := h.UserUsecase.RemoveRelation(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *UserHandler) Check(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		UserName        string `json:"user_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	ok, err := h.UserUsecase.Check(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	if !ok {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return nil
}
