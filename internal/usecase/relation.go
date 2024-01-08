package usecase

import (
	"rbac/utils"

	"github.com/skyrocketOoO/go-utility/set"
	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
	zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"
)

type RelationUsecase struct {
	ZanzibarDagClient *zclient.ZanzibarDagClient
}

func NewRelationUsecase(zanzibarDagClient *zclient.ZanzibarDagClient) *RelationUsecase {
	return &RelationUsecase{
		ZanzibarDagClient: zanzibarDagClient,
	}
}

func (u *RelationUsecase) GetAll() ([]zanzibardagdom.Relation, error) {
	return u.ZanzibarDagClient.GetAll()
}

func (u *RelationUsecase) Query(relation zanzibardagdom.Relation) ([]zanzibardagdom.Relation, error) {
	if err := utils.ValidateReserveWord(relation); err != nil {
		return nil, err
	}
	return u.ZanzibarDagClient.Query(relation)
}

func (u *RelationUsecase) Create(relation zanzibardagdom.Relation) error {
	if err := utils.ValidateReserveWord(relation); err != nil {
		return err
	}
	return u.ZanzibarDagClient.Create(relation, false)
}

func (u *RelationUsecase) Delete(relation zanzibardagdom.Relation) error {
	if err := utils.ValidateReserveWord(relation); err != nil {
		return err
	}
	return u.ZanzibarDagClient.Delete(relation)
}

func (u *RelationUsecase) Check(relation zanzibardagdom.Relation) (bool, error) {
	if err := utils.ValidateReserveWord(relation); err != nil {
		return false, err
	}
	from := zanzibardagdom.Node{
		Namespace: relation.SubjectNamespace,
		Name:      relation.SubjectName,
		Relation:  relation.SubjectRelation,
	}
	to := zanzibardagdom.Node{
		Namespace: relation.ObjectNamespace,
		Name:      relation.ObjectName,
		Relation:  relation.Relation,
	}
	if from.Namespace == "role" && from.Name == "admin" {
		return true, nil
	}

	return u.ZanzibarDagClient.Check(from, to, zanzibardagdom.SearchCondition{})
}

// use to get all relations based on given attr
func (u *RelationUsecase) QueryExistedRelations(namespace, name string) ([]zanzibardagdom.Relation, error) {
	res := set.NewSet[zanzibardagdom.Relation]()

	if name == "" {
		query := zanzibardagdom.Relation{
			ObjectNamespace: namespace,
		}
		relations, err := u.ZanzibarDagClient.Query(query)
		if err != nil {
			return nil, err
		}
		for _, relation := range relations {
			res.Add(relation)
		}

		query = zanzibardagdom.Relation{
			SubjectNamespace: namespace,
		}
		relations, err = u.ZanzibarDagClient.Query(query)
		if err != nil {
			return nil, err
		}
		for _, relation := range relations {
			res.Add(relation)
		}
	} else {
		query := zanzibardagdom.Relation{
			ObjectNamespace: namespace,
			ObjectName:      name,
		}
		relations, err := u.ZanzibarDagClient.Query(query)
		if err != nil {
			return nil, err
		}
		for _, relation := range relations {
			res.Add(relation)
		}

		query = zanzibardagdom.Relation{
			SubjectNamespace: namespace,
			SubjectName:      name,
		}
		relations, err = u.ZanzibarDagClient.Query(query)
		if err != nil {
			return nil, err
		}
		for _, relation := range relations {
			res.Add(relation)
		}
	}

	return res.ToSlice(), nil
}

func (u *RelationUsecase) ClearAllRelations() error {
	return u.ZanzibarDagClient.ClearAllRelations()
}

func (u *RelationUsecase) GetAllObjectRelations(subject zanzibardagdom.Node) ([]zanzibardagdom.Relation, error) {
	return u.ZanzibarDagClient.GetAllObjectRelations(subject, zanzibardagdom.SearchCondition{}, zanzibardagdom.CollectCondition{})
}
