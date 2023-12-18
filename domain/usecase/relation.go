package usecasedomain

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
)

type RelationUsecase interface {
	ListRelations() ([]string, error)
	Link(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error
	Check(relationTuple domain.RelationTuple) (bool, error)
	Path(relationTuple domain.RelationTuple) ([]string, error)
	ListRelationTuples(namespace, name string) ([]sqldomain.RelationTuple, error)
	Create(relationTuple domain.RelationTuple) error
	SafeCreate(relationTuple domain.RelationTuple) error
}
