package delivery

import (
	"rbac/domain"
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

func (h *RelationHandler) ListRelations(c *fiber.Ctx) error {
	relations, err := h.RelationUsecase.ListRelations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(relations)
}

func (h *RelationHandler) Link(c *fiber.Ctx) error {
	type reqBody struct {
		ObjectNamespace     string `json:"object_namespace"`
		ObjectName          string `json:"object_name"`
		Relation            string `json:"relation"`
		SubjectSetNamespace string `json:"subject_set_namespace"`
		SubjectSetName      string `json:"subject_set_name"`
		SubjectSetRelation  string `json:"subject_set_relation"`
	}
	rb := reqBody{}
	if err := c.BodyParser(&rb); err != nil {
		return fiber.NewError(400, "body error")
	}

	err := h.RelationUsecase.Link(rb.ObjectNamespace, rb.ObjectName, rb.Relation, rb.SubjectSetNamespace, rb.SubjectSetName, rb.SubjectSetRelation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Relation linked successfully"})
}

func (h *RelationHandler) Check(c *fiber.Ctx) error {
	type CheckUserRelationReq struct {
		ObjectNamespace     string `json:"object_namespace"`
		ObjectName          string `json:"object_name"`
		Relation            string `json:"relation"`
		SubjectNamespace    string `json:"subject_namespace"`
		SubjectName         string `json:"subject_name"`
		SubjectSetNamespace string `json:"subjectset_namespace"`
		SubjectSetName      string `json:"subjectset_name"`
		SubjectSetRelation  string `json:"subjectset_relation"`
	}
	reqBody := CheckUserRelationReq{}

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(400, "body error")
	}

	ok, err := h.RelationUsecase.Check(domain.RelationTuple{
		ObjectNamespace:           reqBody.ObjectNamespace,
		ObjectName:                reqBody.ObjectName,
		Relation:                  reqBody.Relation,
		SubjectNamespace:          reqBody.SubjectNamespace,
		SubjectName:               reqBody.SubjectName,
		SubjectSetObjectNamespace: reqBody.SubjectSetNamespace,
		SubjectSetObjectName:      reqBody.SubjectSetName,
		SubjectSetRelation:        reqBody.SubjectSetRelation,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if ok {
		return nil
	}
	return c.SendStatus(403)
}

func (h *RelationHandler) Path(c *fiber.Ctx) error {
	type CheckUserRelationReq struct {
		ObjectNamespace     string `json:"object_namespace"`
		ObjectName          string `json:"object_name"`
		Relation            string `json:"relation"`
		SubjectNamespace    string `json:"subject_namespace"`
		SubjectName         string `json:"subject_name"`
		SubjectSetNamespace string `json:"subjectset_namespace"`
		SubjectSetName      string `json:"subjectset_name"`
		SubjectSetRelation  string `json:"subjectset_relation"`
	}
	reqBody := CheckUserRelationReq{}

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(400, "body error")
	}

	path, err := h.RelationUsecase.Path(domain.RelationTuple{
		ObjectNamespace:           reqBody.ObjectNamespace,
		ObjectName:                reqBody.ObjectName,
		Relation:                  reqBody.Relation,
		SubjectNamespace:          reqBody.SubjectNamespace,
		SubjectName:               reqBody.SubjectName,
		SubjectSetObjectNamespace: reqBody.SubjectSetNamespace,
		SubjectSetObjectName:      reqBody.SubjectSetName,
		SubjectSetRelation:        reqBody.SubjectSetRelation,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if len(path) > 0 {
		return c.JSON(path)
	}
	return c.SendStatus(403)
}
