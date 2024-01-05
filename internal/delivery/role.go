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

// @Summary Get all roles
// @Description Get a list of all roles
// @Tags Role
// @Accept json
// @Produce json
// @Success 200 {object} domain.DataResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role [get]
func (h *RoleHandler) GetAll(c *fiber.Ctx) error {
	roles, err := h.RoleUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.StringsResponse{
		Data: roles,
	})
}

// @Summary Delete a role by name
// @Description Delete a role by name
// @Tags Role
// @Accept json
// @Produce json
// @Param name path string true "Role name"
// @Success 200 {string} string "Role deleted successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/{name} [delete]
func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	params := c.AllParams()
	roleName, ok := params["name"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: "get pararm failed",
		})
	}

	err := h.RoleUsecase.Delete(roleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}

// @Summary Add a relation for a role
// @Description Add a relation for a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.AddRelation.request true "Add Relation Request"
// @Success 200 {string} string "Relation added successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/add-relation [post]
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

// @Summary Remove a relation for a role
// @Description Remove a relation for a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.RemoveRelation.request true "Request body to remove a relation for a role"
// @Success 200 {string} string "Relation removed successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/remove-relation [post]
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

// @Summary Add a parent role for a role
// @Description Add a parent role for a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.AddParent.request true "Request body to add a parent role for a role"
// @Success 200 {string} string "Parent role added successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/add-parent [post]
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

// @Summary Remove a parent role for a role
// @Description Remove a parent role for a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.RemoveParent.request true "Request body to remove a parent role for a role"
// @Success 200 {string} string "Parent role removed successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/remove-parent [post]
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

func (h *RoleHandler) GetChildRoles(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	roles, err := h.RoleUsecase.GetChildRoles(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.StringsResponse{
		Data: roles,
	})
}

// @Summary Find all object relations for a role
// @Description Find all object relations for a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.FindAllObjectRelations.request true "Request body to find all object relations for a role"
// @Success 200 {object} domain.DataResponse
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/find-all-object-relations [post]
func (h *RoleHandler) GetAllObjectRelations(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}
	relations, err := h.RoleUsecase.GetAllObjectRelations(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.RelationsResponse{
		Data: relations,
	})
}

// @Summary Get members of a role
// @Description Get members of a role
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.GetMembers.request true "Request body to get members of a role"
// @Success 200 {object} domain.DataResponse
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /role/get-members [post]
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

	return c.JSON(domain.StringsResponse{
		Data: members,
	})
}

// @Summary Check if a role has access to an object
// @Description Check if a role has access to an object
// @Tags Role
// @Accept json
// @Produce json
// @Param request body delivery.Check.request true "Request body to check access for a role to an object"
// @Success 200 {string} string "Role has access to object"
// @Failure 400 {object} domain.ErrResponse
// @Failure 403 {string} string "Role does not have access to object"
// @Failure 500 {object} domain.ErrResponse
// @Router /role/check [post]
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
