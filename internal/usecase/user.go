package usecase

import (
	ucdomain "rbac/domain/usecase"
	"rbac/utils"

	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
	zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"
)

type UserUsecase struct {
	ZanzibarDagClient   *zclient.ZanzibarDagClient
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewUserUsecase(zanzibarDagClient *zclient.ZanzibarDagClient, relationUsecaseRepo ucdomain.RelationUsecase) *UserUsecase {
	return &UserUsecase{
		ZanzibarDagClient:   zanzibarDagClient,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

func (u *UserUsecase) GetAll() ([]string, error) {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelations("user", "")
	if err != nil {
		return nil, err
	}

	users := utils.NewSet[string]()
	for _, tuple := range tuples {
		if tuple.ObjectNamespace == "user" {
			users.Add(tuple.ObjectName)
		}
		if tuple.SubjectNamespace == "user" {
			users.Add(tuple.SubjectName)
		}
	}

	return users.ToSlice(), nil
}

func (u *UserUsecase) Delete(name string) error {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelations("user", name)
	if err != nil {
		return err
	}

	for _, tuple := range tuples {
		if err := u.ZanzibarDagClient.Delete(tuple); err != nil {
			return err
		}
	}
	return nil
}

func (u *UserUsecase) GetRoles(name string) ([]string, error) {
	query := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      name,
	}

	relations, err := u.RelationUsecaseRepo.Query(query)
	if err != nil {
		return nil, err
	}
	roles := utils.NewSet[string]()
	for _, relation := range relations {
		roles.Add(relation.ObjectName)
	}
	return roles.ToSlice(), nil
}

func (u *UserUsecase) AddRole(username, rolename string) error {
	tuple := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       rolename,
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *UserUsecase) RemoveRole(username, rolename string) error {
	tuple := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       rolename,
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      username,
	}
	return u.RelationUsecaseRepo.Delete(tuple)
}

func (u *UserUsecase) GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error) {
	return u.RelationUsecaseRepo.GetAllObjectRelations(
		zanzibardagdom.Node{
			Namespace: "user",
			Name:      name,
		},
	)
}

func (u *UserUsecase) AddRelation(username, relation, objectnamespace, objectname string) error {
	tuple := zanzibardagdom.Relation{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *UserUsecase) RemoveRelation(username, relation, objectnamespace, objectname string) error {
	tuple := zanzibardagdom.Relation{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Delete(tuple)
}

func (u *UserUsecase) Check(userName, relation, objectNamespace, objectName string) (ok bool, err error) {
	return u.RelationUsecaseRepo.Check(zanzibardagdom.Relation{
		ObjectNamespace:  objectNamespace,
		ObjectName:       objectName,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      userName,
	})
}
