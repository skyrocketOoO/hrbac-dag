package usecasedomain

import "rbac/domain"

type RelationUsecase interface {
	ListRelations() ([]string, error)
	Link(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error
	Check(relationTuple domain.RelationTuple) (bool, error)
	Path(relationTuple domain.RelationTuple) ([]string, error)
}
