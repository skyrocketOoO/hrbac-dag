package usecase

import (
	sqldomain "rbac/domain/infra/sql"
)

type RoleUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRoleUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RoleUsecase {
	return &RoleUsecase{
		RelationTupleRepo: relationTupleRepo,
	}
}
