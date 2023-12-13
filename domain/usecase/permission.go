package usecasedomain

import (
	sqldomain "rbac/domain/infra/sql"
)

type PermissionUsecase interface {
	CheckPermission(sqldomain.RelationTuple) (bool, error)
}
