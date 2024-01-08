package usecase

import (
	usecasedomain "rbac/domain/usecase"

	"github.com/skyrocketOoO/go-utility/set"
	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
	zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"
)

type ObjectUsecase struct {
	ZanzibarDagClient *zclient.ZanzibarDagClient
	RoleUsecase       usecasedomain.RoleUsecase
}

func NewObjectUsecase(zanzibarDagClient *zclient.ZanzibarDagClient, roleUsecase usecasedomain.RoleUsecase) *ObjectUsecase {
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
	userRelations := set.NewSet[zanzibardagdom.Relation]()
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
	userRelations := set.NewSet[zanzibardagdom.Relation]()
	for _, r := range relations {
		if r.SubjectNamespace == "role" {
			userRelations.Add(r)
		}
	}

	return userRelations.ToSlice(), nil
}
