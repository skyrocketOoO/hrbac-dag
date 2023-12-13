package usecase

import (
	"errors"
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type UserUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewUserUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *UserUsecase {
	return &UserUsecase{
		RelationTupleRepo: relationTupleRepo,
	}
}

func (uu *UserUsecase) AddUserToRole(username, rolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
		SubNS:    "user",
		SubName:  username,
	}

	return uu.RelationTupleRepo.CreateTuple(tuple)
}

func (uu *UserUsecase) RemoveUserFromRole(username, rolename string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:    "role",
		ObjName:  rolename,
		Relation: "member",
		SubNS:    "user",
		SubName:  username,
	}

	matchedTuples, err := uu.RelationTupleRepo.QueryExactMatchTuples(tuple)
	if err != nil {
		return err
	}
	if len(matchedTuples) == 0 {
		return errors.New("the matched tuples is 0")
	}
	for _, tuple := range matchedTuples {
		if err := uu.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}

	return nil
}

func (uu *UserUsecase) ListUserPermissions(username string) ([]string, error) {
	permissions := utils.NewSet[string]()

	initFilter := sqldomain.RelationTuple{
		SubNS:   "user",
		SubName: username,
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

			tuples, err := uu.RelationTupleRepo.QueryTuples(filter)
			if err != nil {
				return nil, err
			}

			if len(tuples) == 0 && filter.SubNS != "role" {
				// this means it is a leaf object
				permission := filter.SubSetObjNS + ":" + filter.SubSetObjName + "#" + filter.SubSetRelation
				permissions.Add(permission)
			}

			for _, tuple := range tuples {
				reversedTuple := sqldomain.RelationTuple{
					SubSetObjNS:    tuple.ObjNS,
					SubSetObjName:  tuple.ObjName,
					SubSetRelation: tuple.Relation,
				}
				q.Push(reversedTuple)
			}
		}
	}

	return permissions.ToSlice(), nil
}

func (uu *UserUsecase) AddPermissionToUser(username, relation, objectnamespace, objectname string) error {
	tuple := sqldomain.RelationTuple{
		ObjNS:    objectnamespace,
		ObjName:  objectname,
		Relation: relation,
		SubNS:    "user",
		SubName:  username,
	}

	return uu.RelationTupleRepo.CreateTuple(tuple)
}

func (uu *UserUsecase) ListUsers() ([]string, error) {
	filter := sqldomain.RelationTuple{
		SubNS: "user",
	}

	tuples, err := uu.RelationTupleRepo.QueryTuples(filter)
	if err != nil {
		return nil, err
	}

	users := utils.NewSet[string]()
	for _, tuple := range tuples {
		users.Add(tuple.SubName)
	}

	return users.ToSlice(), nil
}
