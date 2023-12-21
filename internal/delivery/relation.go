package delivery

import (
	"fmt"
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

func (h *RelationHandler) GetAllRelations(c *fiber.Ctx) error {
	type response struct {
		data []string
	}
	relations, err := h.RelationUsecase.GetAllRelations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(response{
		data: relations,
	})
}

func (h *RelationHandler) Link(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace  string `json:"object_namespace"`
		ObjectName       string `json:"object_name"`
		Relation         string `json:"relation"`
		SubjectNamespace string `json:"subject_namespace"`
		SubjectName      string `json:"subject_name"`
		SubjectRelation  string `json:"subject_relation"`
	}

	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RelationUsecase.Link(req.ObjectNamespace, req.ObjectName, req.Relation, req.SubjectNamespace, req.SubjectName, req.SubjectRelation)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *RelationHandler) Check(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace  string `json:"object_namespace"`
		ObjectName       string `json:"object_name"`
		Relation         string `json:"relation"`
		SubjectNamespace string `json:"subject_namespace"`
		SubjectName      string `json:"subject_name"`
		SubjectRelation  string `json:"subject_relation"`
	}

	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	ok, err := h.RelationUsecase.Check(domain.RelationTuple{
		ObjectNamespace:  req.ObjectNamespace,
		ObjectName:       req.ObjectName,
		Relation:         req.Relation,
		SubjectNamespace: req.SubjectNamespace,
		SubjectName:      req.SubjectName,
		SubjectRelation:  req.SubjectRelation,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	if ok {
		return nil
	}
	return c.SendStatus(fiber.StatusForbidden)
}

func (h *RelationHandler) Path(c *fiber.Ctx) error {
	type request struct {
		ObjectNamespace  string `json:"object_namespace"`
		ObjectName       string `json:"object_name"`
		Relation         string `json:"relation"`
		SubjectNamespace string `json:"subject_namespace"`
		SubjectName      string `json:"subject_name"`
		SubjectRelation  string `json:"subject_relation"`
	}

	req := request{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	path, err := h.RelationUsecase.GetShortestPath(domain.RelationTuple{
		ObjectNamespace:  req.ObjectNamespace,
		ObjectName:       req.ObjectName,
		Relation:         req.Relation,
		SubjectNamespace: req.SubjectNamespace,
		SubjectName:      req.SubjectName,
		SubjectRelation:  req.SubjectRelation,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	if len(path) > 0 {
		type response struct {
			Data []string
		}
		return c.JSON(response{
			Data: path,
		})
	}
	return c.SendStatus(fiber.StatusForbidden)
}

func (h *RelationHandler) ClearAllRelations(c *fiber.Ctx) error {
	err := h.RelationUsecase.ClearAllRelations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}
