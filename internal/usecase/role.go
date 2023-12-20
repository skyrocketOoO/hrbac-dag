package usecase

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	ucdomain "rbac/domain/usecase"
	"rbac/utils"
)

type RoleUsecase struct {
	RelationTupleRepo   sqldomain.RelationTupleRepository
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewRoleUsecase(relationTupleRepo sqldomain.RelationTupleRepository, relationUsecaseRepo ucdomain.RelationUsecase) *RoleUsecase {
	return &RoleUsecase{
		RelationTupleRepo:   relationTupleRepo,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

func (u *RoleUsecase) ListRoles() ([]string, error) {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("role", "")
	if err != nil {
		return nil, err
	}

	roles := utils.NewSet[string]()
	for _, tuple := range tuples {
		if tuple.ObjectNamespace == "role" {
			roles.Add(tuple.ObjectName)
		}
		if tuple.SubjectNamespace == "role" {
			roles.Add(tuple.SubjectName)
		}
		if tuple.SubjectSetObjectNamespace == "role" {
			roles.Add(tuple.SubjectSetObjectName)
		}
	}

	return roles.ToSlice(), nil
}

func (u *RoleUsecase) GetRole(name string) (string, error) {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("role", name)
	if err != nil {
		return "", err
	}
	if len(tuples) > 0 {
		return name, nil
	}

	return "", nil
}

func (u *RoleUsecase) DeleteRole(name string) error {
	tuples, err := u.RelationUsecaseRepo.QueryExistedRelationTuples("role", name)
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

func (u *RoleUsecase) AddRelation(objnamespace, ObjectName, relation, rolename string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  objnamespace,
		ObjectName:       ObjectName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      rolename,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *RoleUsecase) RemoveRelation(objnamespace, ObjectName, relation, rolename string) error {
	query := domain.RelationTuple{
		ObjectNamespace:  objnamespace,
		ObjectName:       ObjectName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      rolename,
	}

	return u.RelationUsecaseRepo.Delete(query)
}

func (u *RoleUsecase) AddParent(childRolename, parentRolename string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *RoleUsecase) RemoveParent(childRolename, parentRolename string) error {
	query := domain.RelationTuple{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
	}

	return u.RelationUsecaseRepo.Delete(query)
}

func (u *RoleUsecase) FindAllObjectRelations(name string) ([]string, error) {
	return u.RelationUsecaseRepo.FindAllObjectRelations(
		domain.Subject{
			SubjectNamespace: "role",
			SubjectName:      name,
		},
	)
}

func (u *RoleUsecase) GetMembers(name string) ([]string, error) {
	query := domain.RelationTuple{
		ObjectNamespace: "role",
		ObjectName:      name,
		Relation:        "member",
	}

	users := utils.NewSet[string]()

	tuples, err := u.RelationTupleRepo.QueryTuples(query)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		if tuple.SubjectNamespace == "user" {
			users.Add(tuple.SubjectName)
		}
	}
	return users.ToSlice(), nil
}

func (u *RoleUsecase) Check(objectNamespace, objectName, relation, roleName string) (bool, error) {
	return u.RelationUsecaseRepo.Check(domain.RelationTuple{
		ObjectNamespace:  objectNamespace,
		ObjectName:       objectName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      roleName,
	})
}
