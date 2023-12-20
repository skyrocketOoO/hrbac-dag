package usecase

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	ucdomain "rbac/domain/usecase"
	"rbac/utils"
)

type UserUsecase struct {
	RelationTupleRepo   sqldomain.RelationTupleRepository
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewUserUsecase(relationTupleRepo sqldomain.RelationTupleRepository, relationUsecaseRepo ucdomain.RelationUsecase) *UserUsecase {
	return &UserUsecase{
		RelationTupleRepo:   relationTupleRepo,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

func (u *UserUsecase) ListUsers() ([]string, error) {
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
	tuple := domain.RelationTuple{
		ObjectNamespace:  "role",
		ObjectName:       rolename,
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *UserUsecase) RemoveRole(username, rolename string) error {
	tuple := domain.RelationTuple{
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
			SubjectNamespace: "user",
			SubjectName:      name,
		},
	)
}

func (u *UserUsecase) AddRelation(username, relation, objectnamespace, objectname string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *UserUsecase) RemoveRelation(username, relation, objectnamespace, objectname string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Delete(tuple)
}

func (u *UserUsecase) Check(userName, relation, objectNamespace, objectName string) (ok bool, err error) {
	return u.RelationUsecaseRepo.Check(domain.RelationTuple{
		ObjectNamespace:  objectNamespace,
		ObjectName:       objectName,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      userName,
	})
}
