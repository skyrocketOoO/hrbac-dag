package delivery

import (
	"rbac/internal/usecase"
)

type HandlerRepository struct {
	RoleHandler       RoleHandler
	ObjectHandler     ObjectHandler
	UserHandler       UserHandler
	PermissionHandler PermissionHandler
}

func NewHandlerRepository(ucRepo *usecase.UsecaseRepository) *HandlerRepository {
	return &HandlerRepository{
		RoleHandler:       *NewRoleHandler(ucRepo.RoleUsecase),
		ObjectHandler:     *NewObjectHandler(ucRepo.ObjectUsecase),
		UserHandler:       *NewUserHandler(ucRepo.UserUsecase),
		PermissionHandler: *NewPermissionHandler(ucRepo.PermissionUsecase),
	}
}
