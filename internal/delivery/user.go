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

// @Summary Get all users
// @Description Get a list of all users
// @Tags User
// @Produce json
// @Success 200 {object} domain.DataResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/ [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserUsecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: users,
	})
}

// @Summary Get a user by name
// @Description Get user details by name
// @Tags User
// @Produce json
// @Param name path string true "User name"
// @Success 200 {object} domain.DataResponse
// @Failure 400 {object} domain.ErrResponse
// @Failure 404 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/{name} [get]
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

// @Summary Delete a user by name
// @Description Delete a user by name
// @Tags User
// @Produce json
// @Param name path string true "User name"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/{name} [delete]
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

// @Summary Add a role to a user
// @Description Add a role to a user
// @Tags User
// @Accept json
// @Produce json
// @Param request body delivery.AddRole.request true "Request body to add a role to a user"
// @Success 200 {string} string "Role added successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/add-role [post]
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

	err := h.UserUsecase.AddRole(req.UserName, req.RoleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

// @Summary Remove a role from a user
// @Description Remove a role from a user
// @Tags User
// @Accept json
// @Produce json
// @Param request body delivery.RemoveRole.request true "Request body to remove a role from a user"
// @Success 200 {string} string "Role removed successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/remove-role [post]
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

// @Summary Find all object relations for a user
// @Description Find all object relations for a user
// @Tags User
// @Accept json
// @Produce json
// @Param request body delivery.FindAllObjectRelations.request true "Request body to find all object relations for a user"
// @Success 200 {object} domain.DataResponse
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/find-all-object-relations [post]
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

// @Summary Add a relation for a user
// @Description Add a relation for a user
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/add-relation [post]
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

// @Summary Remove a relation for a user
// @Description Remove a relation for a user
// @Tags User
// @Accept json
// @Produce json
// @Param request body delivery.RemoveRelation.request true "Request body to remove a relation for a user"
// @Success 200 {string} string "Relation removed successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /user/remove-relation [post]
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

// @Summary Check if a user has a specific relation for an object
// @Description Check if a user has a specific relation for an object
// @Tags User
// @Accept json
// @Produce json
// @Param request body delivery.Check.request true "Request body to check a relation for a user"
// @Success 200 {string} string "User has the specified relation for the object"
// @Failure 400 {object} domain.ErrResponse
// @Failure 403 {string} string "Forbidden: User does not have the specified relation for the object"
// @Failure 500 {object} domain.ErrResponse
// @Router /user/check [post]
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
