package usecasedomain

import (
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
)

type RelationUsecase interface {
	GetAll() ([]zanzibardagdom.Relation, error)
	Query(relation zanzibardagdom.Relation) ([]zanzibardagdom.Relation, error)
	Create(relation zanzibardagdom.Relation) error
	Delete(relation zanzibardagdom.Relation) error
	Check(relation zanzibardagdom.Relation) (bool, error)
	QueryExistedRelations(namespace, name string) ([]zanzibardagdom.Relation, error)
	ClearAllRelations() error
	GetAllObjectRelations(subject zanzibardagdom.Node) ([]zanzibardagdom.Relation, error)
}
