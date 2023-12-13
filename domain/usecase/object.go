package usecasedomain

import (
	sqldomain "rbac/domain/infra/sql"
)

type ObjectUsecase interface {
	LinkPermission()
	ListPermissionsOnObject(namespace string, name string, relation string) ([]sqldomain.RelationTuple, error)
}
