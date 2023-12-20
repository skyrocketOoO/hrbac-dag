package usecasedomain

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
)

type RelationUsecase interface {
	GetAllRelations() ([]string, error)
	Link(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error
	Check(relationTuple domain.RelationTuple) (bool, error)
	QueryExistedRelationTuples(namespace, name string) ([]sqldomain.RelationTuple, error)
	Create(relationTuple domain.RelationTuple) error
	SafeCreate(relationTuple domain.RelationTuple) error
	ClearAllRelations() error
	FindAllObjectRelations(from domain.Subject) ([]string, error)
	GetShortestPath(relationTuple domain.RelationTuple) ([]string, error)
	GetAllPaths(relationTuple domain.RelationTuple) ([]string, error)
}
