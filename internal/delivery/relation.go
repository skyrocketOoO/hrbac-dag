package delivery

import (
	"fmt"
	"rbac/domain"
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
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
func (h *RelationHandler) GetAll(c *fiber.Ctx) error {
	relations, err := h.RelationUsecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.RelationsResponse{
		Data: relations,
	})
}

func (h *RelationHandler) Query(c *fiber.Ctx) error {
	req := zanzibardagdom.Relation{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	relations, err := h.RelationUsecase.Query(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(domain.RelationsResponse{
		Data: relations,
	})
}

func (h *RelationHandler) Create(c *fiber.Ctx) error {
	req := zanzibardagdom.Relation{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RelationUsecase.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
}

func (h *RelationHandler) Delete(c *fiber.Ctx) error {
	req := zanzibardagdom.Relation{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	err := h.RelationUsecase.Delete(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return nil
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
