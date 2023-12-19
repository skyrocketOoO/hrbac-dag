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
	type reqBody struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	err := h.RoleUsecase.AddRelation(rb.ObjectNamespace, rb.ObjectName, rb.Relation, rb.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission added to role successfully"})
}

func (h *RoleHandler) RemoveRelation(c *fiber.Ctx) error {
	type reqBody struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	err := h.RoleUsecase.RemoveRelation(rb.ObjectNamespace, rb.ObjectName, rb.Relation, rb.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission remove to role successfully"})
}

func (h *RoleHandler) AddParent(c *fiber.Ctx) error {
	type reqBody struct {
		ChildRoleName  string `json:"child_role_name"`
		ParentRoleName string `json:"parent_role_name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	err := h.RoleUsecase.AddParent(rb.ChildRoleName, rb.ParentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Role assigned to role successfully"})
}

func (h *RoleHandler) RemoveParent(c *fiber.Ctx) error {
	type reqBody struct {
		ChildRoleName  string `json:"child_role_name"`
		ParentRoleName string `json:"parent_role_name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	err := h.RoleUsecase.RemoveParent(rb.ChildRoleName, rb.ParentRoleName)
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
	type reqBody struct {
		RoleName string `json:"name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}
	relations, err := h.RoleUsecase.ListRelations(rb.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"relations": relations})
}

func (h *RoleHandler) GetMembers(c *fiber.Ctx) error {
	type reqBody struct {
		RoleName string `json:"name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	members, err := h.RoleUsecase.GetMembers(rb.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"members": members})
}

func (h *RoleHandler) Check(c *fiber.Ctx) error {
	type reqBody struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	ok, err := h.RoleUsecase.Check(rb.ObjectNamespace, rb.ObjectName, rb.Relation, rb.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": ok})
}
