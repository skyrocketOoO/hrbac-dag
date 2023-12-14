package delivery

import (
	usecasedomain "rbac/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	RoleUsecase usecasedomain.RoleUsecase
}

// NewRoleHandler creates a new instance of RoleHandler.
func NewRoleHandler(roleUsecase usecasedomain.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		RoleUsecase: roleUsecase,
	}
}

func (rh *RoleHandler) AddPermissionToRole(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("objnamespace")
	objName := c.FormValue("objname")
	relation := c.FormValue("relation")
	roleName := c.FormValue("rolename")

	// Call the usecase method to add permission to role
	err := rh.RoleUsecase.AddPermissionToRole(objNamespace, objName, relation, roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to role successfully"})
}

func (rh *RoleHandler) AssignRoleUpRole(c *fiber.Ctx) error {
	// Extract data from the request
	childRoleName := c.FormValue("child_rolename")
	parentRoleName := c.FormValue("parent_rolename")

	// Call the usecase method to assign role to role
	err := rh.RoleUsecase.AssignRoleUpRole(childRoleName, parentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role assigned to role successfully"})
}

func (rh *RoleHandler) ListChildRoles(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.Query("rolename")

	// Call the usecase method to list child roles
	childRoles, err := rh.RoleUsecase.ListChildRoles(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"child_roles": childRoles})
}

func (rh *RoleHandler) ListRolePermissions(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.Query("rolename")

	// Call the usecase method to list role permissions
	permissions, err := rh.RoleUsecase.ListRolePermissions(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"permissions": permissions})
}

func (rh *RoleHandler) ListRoles(c *fiber.Ctx) error {
	// Call the usecase method to list all roles
	roles, err := rh.RoleUsecase.ListRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"roles": roles})
}

func (rh *RoleHandler) GetRoleMembers(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.Query("rolename")

	// Call the usecase method to list role members
	members, err := rh.RoleUsecase.GetRoleMembers(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"members": members})
}

func (rh *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.FormValue("rolename")

	// Call the usecase method to delete the role
	err := rh.RoleUsecase.DeleteRole(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}
