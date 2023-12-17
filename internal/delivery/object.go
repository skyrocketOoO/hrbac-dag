package delivery

import (
	usecase "rbac/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type ObjectHandler struct {
	ObjectUsecase usecase.ObjectUsecase
}

func NewObjectHandler(objectUsecase usecase.ObjectUsecase) *ObjectHandler {
	return &ObjectHandler{
		ObjectUsecase: objectUsecase,
	}
}

func (oh *ObjectHandler) ListUserHasRelationOnObject(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")
	relation := c.Query("relation")

	// Call the usecase method to list users with permission
	users, err := oh.ObjectUsecase.ListWhoHasRelationOnObject(objNamespace, objName, relation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}

func (oh *ObjectHandler) ListRoleHasWhatRelationOnObject(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")
	relation := c.Query("relation")

	// Call the usecase method to list roles with permission
	roles, err := oh.ObjectUsecase.ListRolesHasWhatPermissonOnObject(objNamespace, objName, relation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"roles": roles})
}

func (oh *ObjectHandler) ListUserOrRoleHasRelationOnObject(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")
	relation := c.Query("relation")

	// Call the usecase method to list both roles and users with permission
	roles, users, err := oh.ObjectUsecase.ListWhoOrRoleHasRelationOnObject(objNamespace, objName, relation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"roles": roles, "users": users})
}

func (oh *ObjectHandler) ListRelations(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")

	// Call the usecase method to list all permissions
	permissions, err := oh.ObjectUsecase.ListAllRelations(objNamespace, objName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"permissions": permissions})
}
