package usecase

import (
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
	usecasedomain "rbac/domain/usecase"
	"rbac/utils"
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

func (u *ObjectUsecase) GetUserRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {
	relations, err := u.ZanzibarDagClient.GetAllSubjectRelations(
		object,
		zanzibardagdom.SearchCondition{},
		zanzibardagdom.CollectCondition{},
	)
	if err != nil {
		return nil, err
	}
	userRelations := utils.NewSet[zanzibardagdom.Relation]()
	for _, r := range relations {
		if r.SubjectNamespace == "user" {
			userRelations.Add(r)
		}
	}

	return userRelations.ToSlice(), nil
}

func (u *ObjectUsecase) GetRoleRelations(object zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {
	relations, err := u.ZanzibarDagClient.GetAllSubjectRelations(
		object,
		zanzibardagdom.SearchCondition{},
		zanzibardagdom.CollectCondition{},
	)
	if err != nil {
		return nil, err
	}
	userRelations := utils.NewSet[zanzibardagdom.Relation]()
	for _, r := range relations {
		if r.SubjectNamespace == "role" {
			userRelations.Add(r)
		}
	}

	return userRelations.ToSlice(), nil
}
