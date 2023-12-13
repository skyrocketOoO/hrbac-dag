package usecase

import (
	sqldomain "rbac/domain/infra/sql"
	usecasedomain "rbac/domain/usecase"
	"rbac/utils"
)

type ObjectUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
	RoleUsecase       usecasedomain.RoleUsecase
}

func NewObjectUsecase(relationTupleRepo sqldomain.RelationTupleRepository, roleUsecase usecasedomain.RoleUsecase) *ObjectUsecase {
	return &ObjectUsecase{
		RelationTupleRepo: relationTupleRepo,
		RoleUsecase:       roleUsecase,
	}
}

func (ou *ObjectUsecase) LinkPermission(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:          objnamespace,
		ObjName:        objname,
		Relation:       relation,
		SubSetObjNS:    subjnamespace,
		SubSetObjName:  subjname,
		SubSetRelation: subjrelation,
	}

	return ou.RelationTupleRepo.CreateTuple(tuple)
}

func (ou *ObjectUsecase) ListWhoHasRelationOnObject(namespace string, name string, relation string) ([]string, error) {
	users := utils.NewSet[string]()

	initFilter := sqldomain.RelationTuple{
		ObjNS:    namespace,
		ObjName:  name,
		Relation: relation,
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(initFilter)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			filter, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, err
			}

			// if filter.ObjNS == "role" {
			// 	// get all members of role
			// 	members, err := ou.RoleUsecase.GetRoleMembers(filter.ObjName)
			// 	if err != nil {
			// 		return nil, err
			// 	}
			// 	users = append(users, members...)
			// }

			if filter.ObjNS == "user" {
				users.Add(filter.ObjName)
			}

			for _, tuple := range tuples {
				reversedTuple := sqldomain.RelationTuple{
					ObjNS:    tuple.SubSetObjNS,
					ObjName:  tuple.SubSetObjName,
					Relation: tuple.SubSetRelation,
				}
				q.Push(reversedTuple)
			}
		}
	}

	return users.ToSlice(), nil
}

func (ou *ObjectUsecase) ListRolesHasWhatPermissonOnObject(namespace string, name string, relation string) ([]string, error) {
	roles := utils.NewSet[string]()

	initFilter := sqldomain.RelationTuple{
		ObjNS:    namespace,
		ObjName:  name,
		Relation: relation,
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(initFilter)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			filter, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, err
			}

			if filter.ObjNS == "role" {
				roles.Add(filter.ObjName)
			}

			for _, tuple := range tuples {
				reversedTuple := sqldomain.RelationTuple{
					ObjNS:    tuple.SubSetObjNS,
					ObjName:  tuple.SubSetObjName,
					Relation: tuple.SubSetRelation,
				}
				q.Push(reversedTuple)
			}
		}
	}

	return roles.ToSlice(), nil
}

func (ou *ObjectUsecase) ListWhoOrRoleHasWhatPermissionOnObject(namespace string, name string, relation string) ([]string, []string, error) {
	roles := utils.NewSet[string]()
	users := utils.NewSet[string]()

	initFilter := sqldomain.RelationTuple{
		ObjNS:    namespace,
		ObjName:  name,
		Relation: relation,
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(initFilter)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			filter, err := q.Pop()
			if err != nil {
				return nil, nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, nil, err
			}

			if filter.ObjNS == "role" {
				roles.Add(filter.ObjName)
			}
			if filter.ObjNS == "user" {
				users.Add(filter.ObjName)
			}

			for _, tuple := range tuples {
				reversedTuple := sqldomain.RelationTuple{
					ObjNS:    tuple.SubSetObjNS,
					ObjName:  tuple.SubSetObjName,
					Relation: tuple.SubSetRelation,
				}
				q.Push(reversedTuple)
			}
		}
	}

	return roles.ToSlice(), users.ToSlice(), nil
}

func (ou *ObjectUsecase) ListAllPermissions(namespace, name string) ([]string, error) {
	permissions := utils.NewSet[string]()

	filter := sqldomain.RelationTuple{
		ObjNS:   namespace,
		ObjName: name,
	}

	tuples, err := ou.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		permissions.Add(tuple.Relation)
	}

	return permissions.ToSlice(), nil
}
