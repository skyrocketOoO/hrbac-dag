package delivery

import (
	"fmt"
	"rbac/domain"
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
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

func (h *ObjectHandler) GetUserRelations(c *fiber.Ctx) error {
	type requestBody struct {
		Object zanzibardagdom.Node `json:"object"`
	}
	req := requestBody{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	relations, err := h.ObjectUsecase.GetUserRelations(req.Object)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.RelationsResponse{
		Data: relations,
	})
}

func (h *ObjectHandler) GetRoleRelations(c *fiber.Ctx) error {
	type requestBody struct {
		Object zanzibardagdom.Node `json:"object"`
	}
	req := requestBody{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrResponse{
			Error: fmt.Sprintf("Parse body error: %s", err.Error()),
		})
	}

	relations, err := h.ObjectUsecase.GetRoleRelations(req.Object)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(domain.RelationsResponse{
		Data: relations,
	})
}
