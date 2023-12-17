package usecase

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type PermissionUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewPermissionUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *PermissionUsecase {
	return &PermissionUsecase{RelationTupleRepo: relationTupleRepo}
}

func (pu *PermissionUsecase) ListRelations() ([]string, error)

func (ou *ObjectUsecase) Link(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error {
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

func (pu *PermissionUsecase) Check(relationTuple domain.RelationTuple) (bool, error) {
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

func (pu *PermissionUsecase) Path(relationTuple domain.RelationTuple) ([]string, error)
