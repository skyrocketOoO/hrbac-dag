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

func (oh *ObjectHandler) LinkPermission(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("objnamespace")
	objName := c.FormValue("objname")
	relation := c.FormValue("relation")
	subjNamespace := c.FormValue("subjnamespace")
	subjName := c.FormValue("subjname")
	subjRelation := c.FormValue("subjrelation")

	// Call the usecase method to link permission
	err := oh.ObjectUsecase.LinkPermission(objNamespace, objName, relation, subjNamespace, subjName, subjRelation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Permission linked successfully"})
}

func (oh *ObjectHandler) ListWhoHasPermissionOnObject(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")
	relation := c.Query("relation")

	// Call the usecase method to list users with permission
	users, err := oh.ObjectUsecase.ListWhoHasPermissionOnObject(objNamespace, objName, relation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"users": users})
}

func (oh *ObjectHandler) ListRolesHasWhatPermissionOnObject(c *fiber.Ctx) error {
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

func (oh *ObjectHandler) ListWhoOrRoleHasPermissionOnObject(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")
	relation := c.Query("relation")

	// Call the usecase method to list both roles and users with permission
	roles, users, err := oh.ObjectUsecase.ListWhoOrRoleHasPermissionOnObject(objNamespace, objName, relation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"roles": roles, "users": users})
}

func (oh *ObjectHandler) ListAllPermissions(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.Query("objnamespace")
	objName := c.Query("objname")

	// Call the usecase method to list all permissions
	permissions, err := oh.ObjectUsecase.ListAllPermissions(objNamespace, objName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"permissions": permissions})
}
