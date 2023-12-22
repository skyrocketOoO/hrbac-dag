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

// @Summary Get all relations
// @Description Get a list of all relations
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {object} domain.DataResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/get-all-relations [get]
func (h *RelationHandler) GetAllRelations(c *fiber.Ctx) error {
	relations, err := h.RelationUsecase.GetAllRelations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.DataResponse{
		Data: relations,
	})
}

// @Summary Add a relation link
// @Description Add a relation link between two entities
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {string} string "Relation link added successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/add-link [post]
func (h *RelationHandler) AddLink(c *fiber.Ctx) error {
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

	err := h.RelationUsecase.AddLink(
		domain.RelationTuple{
			ObjectNamespace:  req.ObjectNamespace,
			ObjectName:       req.ObjectName,
			Relation:         req.Relation,
			SubjectNamespace: req.SubjectNamespace,
			SubjectName:      req.SubjectName,
			SubjectRelation:  req.SubjectRelation,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}

// @Summary Remove a relation link
// @Description Remove a relation link between two entities
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {string} string "Relation link removed successfully"
// @Failure 400 {object} domain.ErrResponse
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/remove-link [post]
func (h *RelationHandler) RemoveLink(c *fiber.Ctx) error {
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

	err := h.RelationUsecase.RemoveLink(
		domain.RelationTuple{
			ObjectNamespace:  req.ObjectNamespace,
			ObjectName:       req.ObjectName,
			Relation:         req.Relation,
			SubjectNamespace: req.SubjectNamespace,
			SubjectName:      req.SubjectName,
			SubjectRelation:  req.SubjectRelation,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

// @Summary Check if a relation link exists
// @Description Check if a relation link exists between two entities
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {string} string "Relation link exists"
// @Failure 400 {object} domain.ErrResponse
// @Failure 403 {string} string "Relation link does not exist"
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/check [post]
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

// @Summary Get the shortest path between two entities in a relation graph
// @Description Get the shortest path between two entities in a relation graph
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {object} domain.DataResponse "Shortest path between entities"
// @Failure 400 {object} domain.ErrResponse
// @Failure 403 {string} string "No path found"
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/path [post]
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

// @Summary Clear all relations
// @Description Clear all relations in the system
// @Tags Relation
// @Accept json
// @Produce json
// @Success 200 {string} string "All relations cleared"
// @Failure 500 {object} domain.ErrResponse
// @Router /relation/clear-all-relations [post]
func (h *RelationHandler) ClearAllRelations(c *fiber.Ctx) error {
	err := h.RelationUsecase.ClearAllRelations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return nil
}
