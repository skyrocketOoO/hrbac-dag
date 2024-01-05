package usecase

import (
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
	usecasedomain "rbac/domain/usecase"
)

type ObjectUsecase struct {
	ZanzibarDagClient zanzibardagdom.ZanzibarDagRepository
	RoleUsecase       usecasedomain.RoleUsecase
}

func NewObjectUsecase(zanzibarDagClient zanzibardagdom.ZanzibarDagRepository, roleUsecase usecasedomain.RoleUsecase) *ObjectUsecase {
	return &ObjectUsecase{
		ZanzibarDagClient: zanzibarDagClient,
		RoleUsecase:       roleUsecase,
	}
}

func (u *ObjectUsecase) GetUserRelations(namespace string, name string, relation string) ([]string, error)
func (u *ObjectUsecase) GetRoleRelations(namespace string, name string, relation string) ([]string, error)
