package delivery

import (
	"fmt"
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

func (h *RoleHandler) ListRoles(c *fiber.Ctx) error {
	// Call the usecase method to list all roles
	roles, err := h.RoleUsecase.ListRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"roles": roles})
}

func (h *RoleHandler) GetRole(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "parames fault"})
	}
	name, err := h.RoleUsecase.GetRole(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if name == "" {
		err := fmt.Errorf("user %s not found", roleName)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	// Extract data from the request
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "parames fault"})
	}

	// Call the usecase method to delete the role
	err := h.RoleUsecase.DeleteRole(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}

func (h *RoleHandler) AddRelation(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("object_namespace")
	objName := c.FormValue("object_name")
	relation := c.FormValue("relation")
	roleName := c.FormValue("role_name")

	// Call the usecase method to add permission to role
	err := h.RoleUsecase.AddRelation(objNamespace, objName, relation, roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to role successfully"})
}

func (h *RoleHandler) RemoveRelation(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("object_namespace")
	objName := c.FormValue("object_name")
	relation := c.FormValue("relation")
	roleName := c.FormValue("role_name")

	// Call the usecase method to add permission to role
	err := h.RoleUsecase.RemoveRelation(objNamespace, objName, relation, roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission remove to role successfully"})
}

func (h *RoleHandler) AddParent(c *fiber.Ctx) error {
	// Extract data from the request
	childRoleName := c.FormValue("child_role_name")
	parentRoleName := c.FormValue("parent_role_name")

	// Call the usecase method to assign role to role
	err := h.RoleUsecase.AddParent(childRoleName, parentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role assigned to role successfully"})
}

func (h *RoleHandler) RemoveParent(c *fiber.Ctx) error {
	// Extract data from the request
	childRoleName := c.FormValue("child_role_name")
	parentRoleName := c.FormValue("parent_role_name")

	// Call the usecase method to assign role to role
	err := h.RoleUsecase.RemoveParent(childRoleName, parentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role remove to role successfully"})
}

// func (h *RoleHandler) ListChildRoles(c *fiber.Ctx) error {
// 	// Extract data from the request
// 	roleName := c.Query("rolename")

// 	// Call the usecase method to list child roles
// 	childRoles, err := h.RoleUsecase.ListChildRoles(roleName)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(fiber.Map{"child_roles": childRoles})
// }

func (h *RoleHandler) ListRelations(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.Query("name")

	// Call the usecase method to list role relations
	relations, err := h.RoleUsecase.ListRelations(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"relations": relations})
}

func (h *RoleHandler) GetMembers(c *fiber.Ctx) error {
	// Extract data from the request
	roleName := c.Query("name")

	// Call the usecase method to list role members
	members, err := h.RoleUsecase.GetMembers(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"members": members})
}

func (h *RoleHandler) Check(c *fiber.Ctx) error {
	objNs := c.FormValue("object_namespace")
	objName := c.FormValue("object_name")
	relation := c.FormValue("relation")
	roleName := c.FormValue("role_name")

	ok, err := h.RoleUsecase.Check(objNs, objName, relation, roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": ok})
}
