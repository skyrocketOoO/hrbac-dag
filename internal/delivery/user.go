package delivery

import (
	"fmt"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "parames fault"})
	}
	name, err := h.UserUsecase.GetUser(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if name == "" {
		err := fmt.Errorf("user %s not found", roleName)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "parames fault"})
	}
	err := h.UserUsecase.DeleteUser(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
		return fiber.NewError(400, "body error")
	}

	// Call the usecase method to add user to role
	err := h.UserUsecase.AddRole(req.UserName, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User added to role successfully"})
}

func (h *UserHandler) RemoveRole(c *fiber.Ctx) error {
	type request struct {
		UserName string `json:"user_name"`
		RoleName string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "body error")
	}

	// Call the usecase method to remove user from role
	err := h.UserUsecase.RemoveRole(req.UserName, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User removed from role successfully"})
}

func (h *UserHandler) FindAllObjectRelations(c *fiber.Ctx) error {
	// Extract data from the request
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "body error")
	}

	relations, err := h.UserUsecase.FindAllObjectRelations(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"relations": relations})
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
		return fiber.NewError(400, "body error")
	}

	// Call the usecase method to add permission to user
	err := h.UserUsecase.AddRelation(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to user successfully"})
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
		return fiber.NewError(400, "body error")
	}

	// Call the usecase method to add permission to user
	err := h.UserUsecase.RemoveRelation(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to user successfully"})
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
		return fiber.NewError(400, "body error")
	}

	ok, err := h.UserUsecase.Check(req.UserName, req.Relation, req.ObjectNamespace, req.ObjectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": ok})
}
