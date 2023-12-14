package delivery

import (
	usecasedomain "rbac/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	PermissionUsecase usecasedomain.PermissionUsecase
}

func NewPermissionHandler(permissionUsecase usecasedomain.PermissionUsecase) *PermissionHandler {
	return &PermissionHandler{
		PermissionUsecase: permissionUsecase,
	}
}

func (ph *PermissionHandler) CheckUserPermission(c *fiber.Ctx) error {
	type CheckUserPermissionReq struct {
		ObjNS      string `json:"objns"`
		ObjName    string `json:"objname"`
		Permission string `json:"permission"`
		Username   string `json:"username"`
	}
	reqBody := CheckUserPermissionReq{}

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(400, "body error")
	}

	ok, err := ph.PermissionUsecase.CheckUserPermission(
		reqBody.ObjNS,
		reqBody.ObjName,
		reqBody.Permission,
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
