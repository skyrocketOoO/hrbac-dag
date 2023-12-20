package usecase

import (
	"rbac/domain"
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

func (ou *ObjectUsecase) ListUserHasRelationOnObject(namespace string, name string, relation string) ([]string, error) {
	users := utils.NewSet[string]()

	initquery := domain.RelationTuple{
		ObjectNamespace: namespace,
		ObjectName:      name,
		Relation:        relation,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(initquery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, err
			}

			if query.ObjectNamespace == "user" {
				users.Add(query.ObjectName)
			}

			for _, tuple := range tuples {
				nextQuery := domain.RelationTuple{
					ObjectNamespace: tuple.SubjectNamespace,
					ObjectName:      tuple.SubjectName,
					Relation:        tuple.SubjectRelation,
				}
				q.Push(nextQuery)
			}
		}
	}

	return users.ToSlice(), nil
}

func (ou *ObjectUsecase) ListRoleHasWhatRelationOnObject(namespace string, name string, relation string) ([]string, error) {
	roles := utils.NewSet[string]()

	initquery := domain.RelationTuple{
		ObjectNamespace: namespace,
		ObjectName:      name,
		Relation:        relation,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(initquery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, err
			}

			if query.ObjectNamespace == "role" {
				roles.Add(query.ObjectName)
			}

			for _, tuple := range tuples {
				nextQuery := domain.RelationTuple{
					ObjectNamespace: tuple.SubjectNamespace,
					ObjectName:      tuple.SubjectName,
					Relation:        tuple.SubjectRelation,
				}
				q.Push(nextQuery)
			}
		}
	}

	return roles.ToSlice(), nil
}

func (ou *ObjectUsecase) ListUserOrRoleHasRelationOnObject(namespace string, name string, relation string) ([]string, []string, error) {
	roles := utils.NewSet[string]()
	users := utils.NewSet[string]()

	initquery := domain.RelationTuple{
		ObjectNamespace: namespace,
		ObjectName:      name,
		Relation:        relation,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(initquery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return nil, nil, err
			}

			tuples, err := ou.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, nil, err
			}

			if query.ObjectNamespace == "role" {
				roles.Add(query.ObjectName)
			}
			if query.ObjectNamespace == "user" {
				users.Add(query.ObjectName)
			}

			for _, tuple := range tuples {
				nextQuery := domain.RelationTuple{
					ObjectNamespace: tuple.SubjectNamespace,
					ObjectName:      tuple.SubjectName,
					Relation:        tuple.SubjectRelation,
				}
				q.Push(nextQuery)
			}
		}
	}

	return roles.ToSlice(), users.ToSlice(), nil
}

func (ou *ObjectUsecase) ListRelations(namespace, name string) ([]string, error) {
	permissions := utils.NewSet[string]()

	query := domain.RelationTuple{
		ObjectNamespace: namespace,
		ObjectName:      name,
	}

	tuples, err := ou.RelationTupleRepo.QueryTuples(query)
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		permissions.Add(tuple.Relation)
	}

	return permissions.ToSlice(), nil
}
