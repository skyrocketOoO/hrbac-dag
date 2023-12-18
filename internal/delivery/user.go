package delivery

import (
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
	// Call the usecase method to list all users
	users, err := h.UserUsecase.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

}

func (h *UserHandler) AddRole(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("user_name")
	rolename := c.FormValue("role_name")

	// Call the usecase method to add user to role
	err := h.UserUsecase.AddRole(username, rolename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User added to role successfully"})
}

func (h *UserHandler) RemoveRole(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("username")
	rolename := c.FormValue("rolename")

	// Call the usecase method to remove user from role
	err := h.UserUsecase.RemoveRole(username, rolename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User removed from role successfully"})
}

func (h *UserHandler) ListRelations(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.Query("username")

	// Call the usecase method to list user permissions
	permissions, err := h.UserUsecase.ListRelations(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"permissions": permissions})
}

func (h *UserHandler) AddRelation(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("user_name")
	relation := c.FormValue("relation")
	objectNamespace := c.FormValue("object_namespace")
	objectName := c.FormValue("object_name")

	// Call the usecase method to add permission to user
	err := h.UserUsecase.AddRelation(username, relation, objectNamespace, objectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to user successfully"})
}

func (h *UserHandler) RemoveRelation(c *fiber.Ctx) error {

}

func (h *UserHandler) Check(c *fiber.Ctx) error {

}
