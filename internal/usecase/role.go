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

func (ru *RoleUsecase) AddPermissionToRole(objnamespace, objname, relation, rolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:          objnamespace,
		ObjName:        objname,
		Relation:       relation,
		SubSetObjNS:    "role",
		SubSetObjName:  rolename,
		SubSetRelation: "member",
	}

	if err := ru.RelationTupleRepo.CreateTuple(tuple); err != nil {
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

	return ru.RelationTupleRepo.CreateTuple(tuple)
}

func (ru *RoleUsecase) AssignRoleUpRole(childRolename, parentRolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:          "role",
		ObjName:        childRolename,
		Relation:       "parent",
		SubSetObjNS:    "role",
		SubSetObjName:  parentRolename,
		SubSetRelation: "member",
	}

	if err := ru.RelationTupleRepo.CreateTuple(tuple); err != nil {
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

	return ru.RelationTupleRepo.CreateTuple(tuple)
}

func (ru *RoleUsecase) ListChildRoles(rolename string) ([]string, error) {
	initFilter := sqldomain.RelationTuple{
		ObjNS:          "role",
		Relation:       "parent",
		SubSetObjNS:    "role",
		SubSetObjName:  rolename,
		SubSetRelation: "member",
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(initFilter)
	roles := utils.NewSet[string]()

	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			filter, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := ru.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, err
			}

			for _, tuple := range tuples {
				roles.Add(tuple.ObjName)

				reversedTuple := sqldomain.RelationTuple{
					ObjNS:          "role",
					Relation:       "parent",
					SubSetObjNS:    "role",
					SubSetObjName:  tuple.ObjName,
					SubSetRelation: "member",
				}
				q.Push(reversedTuple)
			}
		}
	}

	return roles.ToSlice(), nil
}

func (ru *RoleUsecase) ListRolePermissions(rolename string) ([]string, error) {
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

			tuples, err := ru.RelationTupleRepo.QueryTuples(filter)
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

func (ru *RoleUsecase) ListRoles() ([]string, error) {
	roles := utils.NewSet[string]()

	filter := sqldomain.RelationTuple{
		ObjNS: "role",
	}
	tuples, err := ru.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		roles.Add(tuple.ObjName)
	}

	filter = sqldomain.RelationTuple{
		SubSetObjNS: "role",
	}
	tuples, err = ru.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}
	for _, tuple := range tuples {
		roles.Add(tuple.ObjName)
	}

	return roles.ToSlice(), nil
}

func (ru *RoleUsecase) GetRoleMembers(rolename string) ([]string, error) {
	filter := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
	}

	users := utils.NewSet[string]()

	tuples, err := ru.RelationTupleRepo.QueryTuples(filter)
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

func (ru *RoleUsecase) DeleteRole(rolename string) error {
	filter := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
	}

	tuples, err := ru.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return err
	}
	for _, tuple := range tuples {
		if err := ru.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}

	filter = sqldomain.RelationTuple{
		SubSetObjNS:   "role",
		SubSetObjName: rolename,
	}
	tuples, err = ru.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return err
	}
	for _, tuple := range tuples {
		if err := ru.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}

	return nil
}
