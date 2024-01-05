package usecase

import (
	"rbac/domain"
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
	ucdomain "rbac/domain/usecase"
	"rbac/utils"
)

type UserUsecase struct {
	ZanzibarDagClient   zanzibardagdom.ZanzibarDagRepository
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewUserUsecase(zanzibarDagClient zanzibardagdom.ZanzibarDagRepository, relationUsecaseRepo ucdomain.RelationUsecase) *UserUsecase {
	return &UserUsecase{
		ZanzibarDagClient:   zanzibarDagClient,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

func (u *UserUsecase) GetAllUsers() ([]string, error) {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("user", "")
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
		if tuple.SubjectNamespace == "user" {
			users.Add(tuple.SubjectName)
		}
	}

	return users.ToSlice(), nil
}

// TODO: this method will check existence after list all relation tuples, but we can optimize to first find
func (u *UserUsecase) GetUser(name string) (string, error) {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("user", name)
	if err != nil {
		return "", err
	}

	if len(tuples) > 0 {
		return name, nil
	}
	return "", nil
}

func (u *UserUsecase) DeleteUser(name string) error {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("user", name)
	if err != nil {
		return err
	}

	for _, tuple := range tuples {
		if err := u.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}
	return nil
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

func (u *UserUsecase) FindAllObjectRelations(name string) ([]string, error) {
	return u.RelationUsecaseRepo.FindAllObjectRelations(
		domain.Subject{
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
