package usecasedomain

import zanzibardagdom "rbac/domain/infra/zanzibar-dag"

type ObjectUsecase interface {
	GetUserRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error)
	GetRoleRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error)
}
