package usecasedomain

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
)

type RelationUsecase interface {
	GetAllRelations() ([]string, error)
	AddLink(tuple domain.RelationTuple) error
	RemoveLink(tuple domain.RelationTuple) error
	Check(relationTuple domain.RelationTuple) (bool, error)
	QueryExistedRelationTuples(namespace, name string) ([]sqldomain.RelationTuple, error)
	Create(relationTuple domain.RelationTuple) error
	Delete(relationTuple domain.RelationTuple) error
	ClearAllRelations() error
	FindAllObjectRelations(from domain.Subject) ([]string, error)
	GetShortestPath(relationTuple domain.RelationTuple) ([]string, error)
	GetAllPaths(relationTuple domain.RelationTuple) ([]string, error)
}
