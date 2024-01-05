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
	relationUsecase := NewRelationUsecase(infraRepo.ZanzibarDagRepo)
	roleUsecase := NewRoleUsecase(infraRepo.ZanzibarDagRepo, relationUsecase)

	return &UsecaseRepository{
		ObjectUsecase:   NewObjectUsecase(infraRepo.ZanzibarDagRepo, roleUsecase),
		RelationUsecase: relationUsecase,
		RoleUsecase:     roleUsecase,
		UserUsecase:     NewUserUsecase(infraRepo.ZanzibarDagRepo, relationUsecase),
	}
}
