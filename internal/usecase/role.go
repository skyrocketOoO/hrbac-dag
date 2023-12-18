package usecase

import (
	"fmt"
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
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("role", "")
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
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("role", name)
	if err != nil {
		return "", err
	}
	if len(tuples) > 0 {
		return name, nil
	}

	return "", fmt.Errorf("role %s not found", name)
}

func (u *RoleUsecase) DeleteRole(name string) error {
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("role", name)
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
		ObjectNamespace:           objnamespace,
		ObjectName:                ObjectName,
		Relation:                  relation,
		SubjectSetObjectNamespace: "role",
		SubjectSetObjectName:      rolename,
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *RoleUsecase) RemoveRelation(objnamespace, ObjectName, relation, rolename string) error {
	query := domain.RelationTuple{
		ObjectNamespace:           objnamespace,
		ObjectName:                ObjectName,
		Relation:                  relation,
		SubjectSetObjectNamespace: "role",
		SubjectSetObjectName:      rolename,
	}

	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(query)
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

func (u *RoleUsecase) AddParent(childRolename, parentRolename string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:           "role",
		ObjectName:                childRolename,
		Relation:                  "parent",
		SubjectSetObjectNamespace: "role",
		SubjectSetObjectName:      parentRolename,
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *RoleUsecase) RemoveParent(childRolename, parentRolename string) error {
	query := domain.RelationTuple{
		ObjectNamespace:           "role",
		ObjectName:                childRolename,
		Relation:                  "parent",
		SubjectSetObjectNamespace: "role",
		SubjectSetObjectName:      parentRolename,
	}

	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(query)
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

func (u *RoleUsecase) ListRelations(name string) ([]string, error) {
	initquery := domain.RelationTuple{
		SubjectSetObjectNamespace: "role",
		SubjectSetObjectName:      name,
		SubjectSetRelation:        "member",
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(initquery)
	relations := utils.NewSet[string]()

	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, err
			}

			for _, tuple := range tuples {
				relations.Add(tuple.ObjectNamespace + ":" + tuple.ObjectName + "#" + tuple.Relation)
				if tuple.ObjectNamespace == "role" {
					nextQuery := domain.RelationTuple{
						SubjectSetObjectNamespace: "role",
						SubjectSetObjectName:      tuple.ObjectName,
					}
					q.Push(nextQuery)
				}
				nextQuery := domain.RelationTuple{
					SubjectSetObjectNamespace: tuple.ObjectNamespace,
					SubjectSetObjectName:      tuple.ObjectName,
					SubjectSetRelation:        tuple.Relation,
				}
				q.Push(nextQuery)
			}
		}
	}

	return relations.ToSlice(), nil
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

func (u *RoleUsecase) Check(objectNamespace, objectName, relation, rolename string) (bool, error) {
	firstQuery := domain.RelationTuple{
		SubjectNamespace: "role",
		SubjectName:      rolename,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return false, err
			}

			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				if tuple.ObjectName == objectName && tuple.ObjectNamespace == objectNamespace && tuple.Relation == relation {
					return true, nil
				}

				if tuple.ObjectNamespace == "role" {
					// use role to search
					nextQuery := domain.RelationTuple{
						SubjectNamespace: "role",
						SubjectName:      tuple.ObjectName,
					}
					q.Push(nextQuery)
				}

				nextQuery := domain.RelationTuple{
					SubjectSetObjectNamespace: tuple.ObjectNamespace,
					SubjectSetObjectName:      tuple.ObjectName,
					SubjectSetRelation:        tuple.Relation,
				}
				q.Push(nextQuery)
			}
		}
	}
	return false, nil
}
