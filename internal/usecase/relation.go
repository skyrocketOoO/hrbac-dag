package usecase

import (
	"errors"
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	"rbac/utils"
)

type RelationUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRelationUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RelationUsecase {
	return &RelationUsecase{
		RelationTupleRepo: relationTupleRepo,
	}
}

func (u *RelationUsecase) ListRelations() ([]string, error) {
	tuples, err := u.RelationTupleRepo.GetTuples()
	if err != nil {
		return nil, err
	}

	relations := []string{}
	for _, tuple := range tuples {
		relations = append(relations, utils.RelationTupleToString(utils.ConvertRelationTuple(tuple)))
	}

	return relations, nil
}

func (u *RelationUsecase) Create(relationTuple domain.RelationTuple) error {
	if relationTuple.ObjectNamespace != "role" && relationTuple.ObjectNamespace != "user" {
		// object case: use <ns>:<name>#<relation> as instance
		if relationTuple.ObjectName != "*" && relationTuple.Relation != "*" {
			parentTuple := domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                relationTuple.ObjectName,
				Relation:                  relationTuple.Relation,
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      "*",
				SubjectSetRelation:        relationTuple.Relation,
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}

			parentTuple = domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                relationTuple.ObjectName,
				Relation:                  relationTuple.Relation,
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      relationTuple.ObjectName,
				SubjectSetRelation:        "*",
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}

			parentTuple = domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                "*",
				Relation:                  relationTuple.Relation,
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      "*",
				SubjectSetRelation:        "*",
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}

			parentTuple = domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                relationTuple.ObjectName,
				Relation:                  "*",
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      "*",
				SubjectSetRelation:        "*",
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}

		} else if relationTuple.ObjectName != "*" {
			parentTuple := domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                "*",
				Relation:                  relationTuple.Relation,
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      "*",
				SubjectSetRelation:        "*",
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}
		} else if relationTuple.Relation != "*" {
			parentTuple := domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                relationTuple.ObjectName,
				Relation:                  "*",
				SubjectSetObjectNamespace: relationTuple.ObjectNamespace,
				SubjectSetObjectName:      "*",
				SubjectSetRelation:        "*",
			}
			if err := u.SafeCreate(parentTuple); err != nil {
				return err
			}
		}
	}

	return u.SafeCreate(relationTuple)
}

func (u *RelationUsecase) SafeCreate(relationTuple domain.RelationTuple) error {
	// if tuple exist, return
	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(relationTuple)
	if err != nil {
		return err
	}
	if len(tuples) > 0 {
		return nil
	}
	return u.RelationTupleRepo.CreateTuple(relationTuple)
}

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
	var firstQuery domain.RelationTuple
	if relationTuple.SubjectNamespace != "" {
		firstQuery.SubjectNamespace = relationTuple.SubjectNamespace
		firstQuery.SubjectName = relationTuple.SubjectName
	} else {
		firstQuery.SubjectSetObjectNamespace = relationTuple.SubjectSetObjectNamespace
		firstQuery.SubjectSetObjectName = relationTuple.SubjectSetObjectName
		firstQuery.SubjectSetRelation = relationTuple.SubjectSetRelation
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
				if tuple.ObjectNamespace == relationTuple.ObjectNamespace && tuple.ObjectName == relationTuple.ObjectName && tuple.Relation == relationTuple.Relation {
					return true, nil
				}
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

	return false, nil
}

// TODO: use bfs to return shortest path or return all paths
func (u *RelationUsecase) Path(relationTuple domain.RelationTuple) ([]string, error) {
	paths := []string{}

	var dfs func(relationTuple domain.RelationTuple, curPath []string) error
	dfs = func(relationTuple domain.RelationTuple, curPath []string) error {
		if len(paths) > 0 {
			return nil
		}

		query := domain.RelationTuple{}
		if relationTuple.SubjectNamespace != "" {
			query.SubjectNamespace = relationTuple.SubjectNamespace
			query.SubjectName = relationTuple.SubjectName
		} else {
			query.SubjectSetObjectNamespace = relationTuple.SubjectSetObjectNamespace
			query.SubjectSetObjectName = relationTuple.SubjectSetObjectName
			query.SubjectSetRelation = relationTuple.SubjectSetRelation
		}

		tuples, err := u.RelationTupleRepo.QueryTuples(query)
		if err != nil {
			return err
		}

		for _, tuple := range tuples {
			if tuple.ObjectName == relationTuple.ObjectName &&
				tuple.ObjectNamespace == relationTuple.ObjectNamespace &&
				tuple.Relation == relationTuple.Relation {
				paths = append(paths, curPath...)
				paths = append(paths, utils.RelationTupleToString(relationTuple))
				return nil
			}
			curPath = append(curPath, utils.RelationTupleToString(utils.ConvertRelationTuple(tuple)))

			if tuple.ObjectNamespace == "role" {
				nextQuery := domain.RelationTuple{
					ObjectNamespace:           relationTuple.ObjectNamespace,
					ObjectName:                relationTuple.ObjectName,
					Relation:                  relationTuple.Relation,
					SubjectSetObjectNamespace: "role",
					SubjectSetObjectName:      tuple.ObjectName,
				}
				if err := dfs(nextQuery, curPath); err != nil {
					return err
				}
			}
			nextQuery := domain.RelationTuple{
				ObjectNamespace:           relationTuple.ObjectNamespace,
				ObjectName:                relationTuple.ObjectName,
				Relation:                  relationTuple.Relation,
				SubjectSetObjectNamespace: tuple.ObjectNamespace,
				SubjectSetObjectName:      tuple.ObjectName,
				SubjectSetRelation:        tuple.Relation,
			}
			if err := dfs(nextQuery, curPath); err != nil {
				return err
			}

			curPath = curPath[:len(curPath)-1]
		}
		return nil
	}

	emptyPath := []string{}
	if err := dfs(relationTuple, emptyPath); err != nil {
		return nil, err
	}
	if len(paths) > 0 {
		return paths, nil
	}
	return nil, errors.New("paths not exist")
}

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
