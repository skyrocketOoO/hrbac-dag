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

func (uh *UserHandler) AddUserToRole(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("username")
	rolename := c.FormValue("rolename")

	// Call the usecase method to add user to role
	err := uh.UserUsecase.AddUserToRole(username, rolename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User added to role successfully"})
}

func (uh *UserHandler) RemoveUserFromRole(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("username")
	rolename := c.FormValue("rolename")

	// Call the usecase method to remove user from role
	err := uh.UserUsecase.RemoveUserFromRole(username, rolename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User removed from role successfully"})
}

func (uh *UserHandler) ListUserPermissions(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.Query("username")

	// Call the usecase method to list user permissions
	permissions, err := uh.UserUsecase.ListUserPermissions(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"permissions": permissions})
}

func (uh *UserHandler) AddPermissionToUser(c *fiber.Ctx) error {
	// Extract data from the request
	username := c.FormValue("username")
	relation := c.FormValue("relation")
	objectNamespace := c.FormValue("objectnamespace")
	objectName := c.FormValue("objectname")

	// Call the usecase method to add permission to user
	err := uh.UserUsecase.AddPermissionToUser(username, relation, objectNamespace, objectName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to user successfully"})
}

func (uh *UserHandler) ListUsers(c *fiber.Ctx) error {
	// Call the usecase method to list all users
	users, err := uh.UserUsecase.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}
