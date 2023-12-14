package usecase

import (
	ucdomain "rbac/domain/usecase"
	"rbac/internal/infra/sql"
)

type UsecaseRepository struct {
	ObjectUsecase     ucdomain.ObjectUsecase
	PermissionUsecase ucdomain.PermissionUsecase
	RoleUsecase       ucdomain.RoleUsecase
	UserUsecase       ucdomain.UserUsecase
}

func NewUsecaseRepository(sqlRepo *sql.OrmRepository) *UsecaseRepository {
	roleUsecase := NewRoleUsecase(&sqlRepo.RelationshipRepo)

	return &UsecaseRepository{
		ObjectUsecase:     NewObjectUsecase(&sqlRepo.RelationshipRepo, roleUsecase),
		PermissionUsecase: NewPermissionUsecase(&sqlRepo.RelationshipRepo),
		RoleUsecase: roleUsecase,
		UserUsecase: NewUserUsecase(&sqlRepo.RelationshipRepo),
	}
}
