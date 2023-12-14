package usecase

import (
	sqldomain "rbac/domain/infra/sql"
)

type PermissionUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewPermissionUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *PermissionUsecase {
	return &PermissionUsecase{RelationTupleRepo: relationTupleRepo}
}

func (pu *PermissionUsecase) CheckPermission(sqldomain.RelationTuple) (bool, error) {

	return false, nil
}
