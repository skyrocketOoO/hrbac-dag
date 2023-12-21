package delivery

import (
	"fmt"
	"rbac/domain"
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

func (h *RoleHandler) GetAllRoles(c *fiber.Ctx) error {
	roles, err := h.RoleUsecase.GetAllRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: roles,
	})
}

func (h *RoleHandler) GetRole(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: "get pararm failed",
		})
	}

	name, err := h.RoleUsecase.GetRole(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	if name == "" {
		return c.Status(fiber.StatusNotFound).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("user %s not found", roleName),
		})
	}
	return nil
}

func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: "get pararm failed",
		})
	}

	err := h.RoleUsecase.DeleteRole(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}

func (h *RoleHandler) AddRelation(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RoleUsecase.AddRelation(req.ObjectNamespace, req.ObjectName, req.Relation, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *RoleHandler) RemoveRelation(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RoleUsecase.RemoveRelation(req.ObjectNamespace, req.ObjectName, req.Relation, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *RoleHandler) AddParent(c *fiber.Ctx) error {
	type request struct {
		ChildRoleName  string `json:"child_role_name"`
		ParentRoleName string `json:"parent_role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RoleUsecase.AddParent(req.ChildRoleName, req.ParentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *RoleHandler) RemoveParent(c *fiber.Ctx) error {
	type request struct {
		ChildRoleName  string `json:"child_role_name"`
		ParentRoleName string `json:"parent_role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RoleUsecase.RemoveParent(req.ChildRoleName, req.ParentRoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

// func (h *RoleHandler) GetAllChildRoles(c *fiber.Ctx) error {
// 	// Extract data from the request
// 	roleName := c.Query("rolename")

// 	// Call the usecase method to list child roles
// 	childRoles, err := h.RoleUsecase.GetAllChildRoles(roleName)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(fiber.Map{"child_roles": childRoles})
// }

func (h *RoleHandler) FindAllObjectRelations(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}
	relations, err := h.RoleUsecase.FindAllObjectRelations(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: relations,
	})
}

func (h *RoleHandler) GetMembers(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	members, err := h.RoleUsecase.GetMembers(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: members,
	})
}

func (h *RoleHandler) Check(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace string `json:"object_namespace"`
		ObjectName      string `json:"object_name"`
		Relation        string `json:"relation"`
		RoleName        string `json:"role_name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	ok, err := h.RoleUsecase.Check(req.ObjectNamespace, req.ObjectName, req.Relation, req.RoleName)
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
