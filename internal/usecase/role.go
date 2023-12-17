package usecase

import (
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type RoleUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRoleUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RoleUsecase {
	return &RoleUsecase{
		RelationTupleRepo: relationTupleRepo,
	}
}

func (u *RoleUsecase) ListRoles() ([]string, error) {
	roles := utils.NewSet[string]()

	filter := sqldomain.RelationTuple{
		ObjNS: "role",
	}
	tuples, err := u.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		roles.Add(tuple.ObjName)
	}

	filter = sqldomain.RelationTuple{
		SubSetObjNS: "role",
	}
	tuples, err = u.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		roles.Add(tuple.ObjName)
	}

	return roles.ToSlice(), nil
}

func (u *RoleUsecase) GetRole(name string) (string, error)

func (u *RoleUsecase) DeleteRole(rolename string) error {
	filter := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
	}

	tuples, err := u.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return err
	}
	for _, tuple := range tuples {
		if err := u.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}

	filter = sqldomain.RelationTuple{
		SubSetObjNS:   "role",
		SubSetObjName: rolename,
	}
	tuples, err = u.RelationTupleRepo.QueryTuples(filter)
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

func (u *RoleUsecase) AddRelation(objnamespace, objname, relation, rolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:          objnamespace,
		ObjName:        objname,
		Relation:       relation,
		SubSetObjNS:    "role",
		SubSetObjName:  rolename,
		SubSetRelation: "member",
	}

	if err := u.RelationTupleRepo.CreateTuple(tuple); err != nil {
		return err
	}

	tuple = sqldomain.RelationTuple{
		ObjNS:          objnamespace,
		ObjName:        objname,
		Relation:       relation,
		SubSetObjNS:    "role",
		SubSetObjName:  rolename,
		SubSetRelation: "parent",
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *RoleUsecase) RemoveRelation(objnamespace, objname, relation, rolename string) error

func (u *RoleUsecase) AddParent(childRolename, parentRolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:          "role",
		ObjName:        childRolename,
		Relation:       "parent",
		SubSetObjNS:    "role",
		SubSetObjName:  parentRolename,
		SubSetRelation: "member",
	}

	if err := u.RelationTupleRepo.CreateTuple(tuple); err != nil {
		return err
	}

	tuple = sqldomain.RelationTuple{
		ObjNS:          "role",
		ObjName:        childRolename,
		Relation:       "parent",
		SubSetObjNS:    "role",
		SubSetObjName:  parentRolename,
		SubSetRelation: "parent",
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *RoleUsecase) RemoveParent(childRolename, parentRolename string) error

func (u *RoleUsecase) ListRelations(name string) ([]string, error) {
	initFilter := sqldomain.RelationTuple{
		SubSetObjNS:    "role",
		SubSetObjName:  rolename,
		SubSetRelation: "member",
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(initFilter)
	permissions := utils.NewSet[string]()

	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			filter, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := u.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, err
			}

			for _, tuple := range tuples {
				if tuple.ObjNS == "role" {
					reversedTuple := sqldomain.RelationTuple{
						SubSetObjNS:    "role",
						SubSetObjName:  tuple.ObjName,
						SubSetRelation: "member",
					}
					q.Push(reversedTuple)
				} else {
					// permission
					permissions.Add(tuple.ObjNS + ":" + tuple.ObjName + "#" + tuple.Relation)
					reversedTuple := sqldomain.RelationTuple{
						SubSetObjNS:    tuple.ObjNS,
						SubSetObjName:  tuple.ObjName,
						SubSetRelation: tuple.Relation,
					}
					q.Push(reversedTuple)
				}
			}
		}
	}

	return permissions.ToSlice(), nil
}

func (u *RoleUsecase) GetMembers(name string) ([]string, error) {
	filter := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
	}

	users := utils.NewSet[string]()

	tuples, err := u.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		if tuple.SubNS == "user" {
			users.Add(tuple.SubName)
		}
	}
	return users.ToSlice(), nil
}

func (u *RoleUsecase) Check(objnamespace, objname, relation, rolename string) (bool, error)
