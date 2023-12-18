package usecase

import (
	ucdomain "rbac/domain/usecase"
	"rbac/internal/infra/sql"
)

type UsecaseRepository struct {
	ObjectUsecase   ucdomain.ObjectUsecase
	RelationUsecase ucdomain.RelationUsecase
	RoleUsecase     ucdomain.RoleUsecase
	UserUsecase     ucdomain.UserUsecase
}

func NewUsecaseRepository(sqlRepo *sql.OrmRepository) *UsecaseRepository {
	relationUsecase := NewRelationUsecase(&sqlRepo.RelationshipRepo)
	roleUsecase := NewRoleUsecase(&sqlRepo.RelationshipRepo, relationUsecase)

	return &UsecaseRepository{
		ObjectUsecase:   NewObjectUsecase(&sqlRepo.RelationshipRepo, roleUsecase),
		RelationUsecase: relationUsecase,
		RoleUsecase:     roleUsecase,
		UserUsecase:     NewUserUsecase(&sqlRepo.RelationshipRepo, relationUsecase),
	}
}
