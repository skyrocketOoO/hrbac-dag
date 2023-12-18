package usecase

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type RelationUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRelationUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RelationUsecase {
	return &RelationUsecase{RelationTupleRepo: relationTupleRepo}
}

func (u *RelationUsecase) ListRelations() ([]string, error)

func (u *RelationUsecase) Link(objnamespace, ObjectName, relation, subjnamespace, subjname, subjrelation string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:           objnamespace,
		ObjectName:                ObjectName,
		Relation:                  relation,
		SubjectSetObjectNamespace: subjnamespace,
		SubjectSetObjectName:      subjname,
		SubjectSetRelation:        subjrelation,
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *RelationUsecase) Check(relationTuple domain.RelationTuple) (bool, error) {
	firstQuery := domain.RelationTuple{
		SubjectNamespace: "user",
		SubjectName:      relationTuple.SubjectName,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				if tuple.ObjectNamespace == relationTuple.ObjectNamespace && tuple.ObjectName == relationTuple.ObjectName && tuple.Relation == relationTuple.Relation {
					return true, nil
				}

				newQuery := domain.RelationTuple{
					SubjectSetObjectNamespace: tuple.ObjectNamespace,
					SubjectSetObjectName:      tuple.ObjectName,
					SubjectSetRelation:        tuple.Relation,
				}
				q.Push(newQuery)
			}
		}
	}

	return false, nil
}

func (u *RelationUsecase) Path(relationTuple domain.RelationTuple) ([]string, error)

// use to get all relations based on given attr
func (u *RelationUsecase) ListRelationTuples(namespace, name string) ([]sqldomain.RelationTuple, error) {
	res := utils.NewSet[sqldomain.RelationTuple]()

	if name == "" {
		query := domain.RelationTuple{
			ObjectNamespace: namespace,
		}
		tuples, err := u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}

		query = domain.RelationTuple{
			SubjectSetObjectNamespace: namespace,
		}
		tuples, err = u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}

		query = domain.RelationTuple{
			SubjectNamespace: namespace,
		}
		tuples, err = u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}
	} else {
		query := domain.RelationTuple{
			ObjectNamespace: namespace,
			ObjectName:      name,
		}
		tuples, err := u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}

		query = domain.RelationTuple{
			SubjectSetObjectNamespace: namespace,
			SubjectSetObjectName:      name,
		}
		tuples, err = u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}

		query = domain.RelationTuple{
			SubjectNamespace: namespace,
			SubjectName:      name,
		}
		tuples, err = u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return nil, err
		}
		for _, tuple := range tuples {
			res.Add(tuple)
		}
	}

	return res.ToSlice(), nil
}
