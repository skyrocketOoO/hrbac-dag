package delivery

import (
	usecasedomain "rbac/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type RelationHandler struct {
	RelationUsecase usecasedomain.RelationUsecase
}

func NewRelationHandler(permissionUsecase usecasedomain.RelationUsecase) *RelationHandler {
	return &RelationHandler{
		RelationUsecase: permissionUsecase,
	}
}

func (oh *RelationHandler) ListRelations(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("objnamespace")
	objName := c.FormValue("objname")
	relation := c.FormValue("relation")
	subjNamespace := c.FormValue("subjnamespace")
	subjName := c.FormValue("subjname")
	subjRelation := c.FormValue("subjrelation")

	// Call the usecase method to link permission
	err := oh.ObjectUsecase.LinkRelation(objNamespace, objName, relation, subjNamespace, subjName, subjRelation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Relation linked successfully"})
}

func (ph *RelationHandler) Link(c *fiber.Ctx) error {
	// Extract data from the request
	objNamespace := c.FormValue("objnamespace")
	objName := c.FormValue("objname")
	relation := c.FormValue("relation")
	subjNamespace := c.FormValue("subjnamespace")
	subjName := c.FormValue("subjname")
	subjRelation := c.FormValue("subjrelation")

	// Call the usecase method to link permission
	err := oh.ObjectUsecase.LinkRelation(objNamespace, objName, relation, subjNamespace, subjName, subjRelation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Relation linked successfully"})
}

func (ph *RelationHandler) Check(c *fiber.Ctx) error {
	type CheckUserRelationReq struct {
		ObjNS    string `json:"obj_ns"`
		ObjName  string `json:"obj_name"`
		Relation string `json:"relation"`
		Username string `json:"user_name"`
	}
	reqBody := CheckUserRelationReq{}

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(400, "body error")
	}

	ok, err := ph.RelationUsecase.CheckUserRelation(
		reqBody.ObjNS,
		reqBody.ObjName,
		reqBody.Relation,
		reqBody.Username,
	)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return c.SendStatus(403)
}

func (ph *RelationHandler) Path(c *fiber.Ctx) error
