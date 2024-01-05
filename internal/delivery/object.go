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

func (h *ObjectHandler) GetUserRelations(c *fiber.Ctx) error

func (h *ObjectHandler) GetRoleRelations(c *fiber.Ctx) error
