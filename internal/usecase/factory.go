package usecase

import (
	ucdomain "rbac/domain/usecase"
	"rbac/internal/infra"
)

type UsecaseRepository struct {
	ObjectUsecase   ucdomain.ObjectUsecase
	RelationUsecase ucdomain.RelationUsecase
	RoleUsecase     ucdomain.RoleUsecase
	UserUsecase     ucdomain.UserUsecase
}

func NewUsecaseRepository(infraRepo *infra.InfraRepository) *UsecaseRepository {
	relationUsecase := NewRelationUsecase(infraRepo.ZanzibarDagClient)
	roleUsecase := NewRoleUsecase(infraRepo.ZanzibarDagClient, relationUsecase)

	return &UsecaseRepository{
		ObjectUsecase:   NewObjectUsecase(infraRepo.ZanzibarDagClient, roleUsecase),
		RelationUsecase: relationUsecase,
		RoleUsecase:     roleUsecase,
		UserUsecase:     NewUserUsecase(infraRepo.ZanzibarDagClient, relationUsecase),
	}
}
