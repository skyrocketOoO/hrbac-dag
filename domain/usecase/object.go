package usecasedomain

import zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"

type ObjectUsecase interface {
	GetUserRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error)
	GetRoleRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error)
}
