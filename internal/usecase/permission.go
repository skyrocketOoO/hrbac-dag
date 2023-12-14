package usecase

import (
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type PermissionUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewPermissionUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *PermissionUsecase {
	return &PermissionUsecase{RelationTupleRepo: relationTupleRepo}
}

func (pu *PermissionUsecase) CheckUserPermission(objNS, objName, Permission, username string) (bool, error) {
	query := sqldomain.RelationTuple{
		SubNS:   "user",
		SubName: username,
	}

	q := utils.NewQueue[sqldomain.RelationTuple]()
	q.Push(query)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			tuples, err := pu.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				if tuple.ObjNS == objNS && tuple.ObjName == objName && tuple.Relation == Permission {
					return true, nil
				}

				newQuery := sqldomain.RelationTuple{
					SubSetObjNS:    tuple.ObjNS,
					SubSetObjName:  tuple.ObjName,
					SubSetRelation: tuple.Relation,
				}
				q.Push(newQuery)
			}
		}
	}

	return false, nil
}
