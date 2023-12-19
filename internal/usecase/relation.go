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
	// fmt.Printf("Person: %+v\n", relationTuple)
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
	return u.searchTemplate(
		domain.Subject{
			SubjectNamespace:    relationTuple.SubjectNamespace,
			SubjectName:         relationTuple.SubjectName,
			SubjectSetNamespace: relationTuple.SubjectSetObjectNamespace,
			SubjectSetName:      relationTuple.SubjectSetObjectName,
			SubjectSetRelation:  relationTuple.SubjectSetRelation,
		},
		domain.Object{
			ObjectNamespace: relationTuple.ObjectNamespace,
			ObjectName:      relationTuple.ObjectName,
			Relation:        relationTuple.Relation,
		},
	)
}

// TODO: use bfs to return shortest path or return all paths
func (u *RelationUsecase) Path(relationTuple domain.RelationTuple) ([]string, error) {

	var dfs func(relationTuple domain.RelationTuple, curPath []string, finalPath *[]string) error
	dfs = func(relationTuple domain.RelationTuple, curPath []string, finalPath *[]string) error {
		if len(*finalPath) > 0 {
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
				*finalPath = append(*finalPath, curPath...)
				*finalPath = append(*finalPath, utils.RelationTupleToString(relationTuple))
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
				if err := dfs(nextQuery, curPath, finalPath); err != nil {
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
			if err := dfs(nextQuery, curPath, finalPath); err != nil {
				return err
			}

			curPath = curPath[:len(curPath)-1]
		}
		return nil
	}

	paths := []string{}
	emptyPath := []string{}
	if err := dfs(relationTuple, emptyPath, &paths); err != nil {
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

func (u *RelationUsecase) ClearAllRelations() error {
	return u.RelationTupleRepo.DeleteAllTuples()
}

func (u *RelationUsecase) searchTemplate(from domain.Subject, to domain.Object) (ok bool, err error) {
	// usecase
	// (u *RoleUsecase) ListRelations(namespace, name string) ([]string, error)
	// (u *RoleUsecase) Check(objectNamespace, objectName, relation, rolename string) (bool, error)
	// Path(relationTuple domain.RelationTuple) ([]string, error)

	var firstQuery domain.RelationTuple
	if from.SubjectNamespace != "" {
		firstQuery.SubjectNamespace = from.SubjectNamespace
		firstQuery.SubjectName = from.SubjectName
	} else if from.SubjectSetNamespace != "" {
		firstQuery.SubjectSetObjectNamespace = from.SubjectSetNamespace
		firstQuery.SubjectSetObjectName = from.SubjectSetName
		firstQuery.SubjectSetRelation = from.SubjectSetRelation
	} else {
		return false, errors.New("subject error")
	}
	firstQuery.ObjectNamespace = to.ObjectNamespace
	firstQuery.ObjectName = to.ObjectName
	firstQuery.Relation = to.Relation

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, _ := q.Pop()
			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				if tuple.ObjectNamespace == to.ObjectNamespace && tuple.ObjectName == to.ObjectName && tuple.Relation == to.Relation {
					return true, nil
				}
				// object case
				// TODO: the <ns> <*> <*> query will direct find the next layer object, so can directly jump to next
				if tuple.ObjectNamespace != "role" && tuple.ObjectNamespace != "user" {
					if tuple.ObjectName == "*" && tuple.Relation == "*" {
						if tuple.ObjectNamespace == to.ObjectNamespace {
							return true, nil
						}
						// add abstract link
						nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
							SubjectSetObjectNamespace: tuple.ObjectNamespace,
							SubjectSetObjectName:      "*",
						})
						if err != nil {
							return false, err
						}
						set := utils.NewSet[string]()
						for _, nq := range nextQueries {
							set.Add(nq.SubjectSetRelation)
						}
						for _, rel := range set.ToSlice() {
							q.Push(domain.RelationTuple{
								SubjectSetObjectNamespace: tuple.ObjectNamespace,
								SubjectSetObjectName:      "*",
								SubjectSetRelation:        rel,
							})
						}

						nextQueries, err = u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
							SubjectSetObjectNamespace: tuple.ObjectNamespace,
							SubjectSetRelation:        "*",
						})
						if err != nil {
							return false, err
						}
						set = utils.NewSet[string]()
						for _, nq := range nextQueries {
							set.Add(nq.SubjectSetObjectName)
						}
						for _, name := range set.ToSlice() {
							q.Push(domain.RelationTuple{
								SubjectSetObjectNamespace: tuple.ObjectNamespace,
								SubjectSetObjectName:      name,
								SubjectSetRelation:        "*",
							})
						}
					} else if tuple.ObjectName == "*" {
						if tuple.ObjectNamespace == to.ObjectNamespace && tuple.Relation == to.Relation {
							return true, nil
						}
						// abstract link
						nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
							SubjectSetObjectNamespace: tuple.ObjectNamespace,
							SubjectSetRelation:        tuple.Relation,
						})
						if err != nil {
							return false, err
						}
						set := utils.NewSet[string]()
						for _, nq := range nextQueries {
							set.Add(nq.ObjectName)
						}
						for _, name := range set.ToSlice() {
							q.Push(domain.RelationTuple{
								SubjectSetObjectNamespace: tuple.ObjectNamespace,
								SubjectSetObjectName:      name,
								SubjectSetRelation:        tuple.Relation,
							})
						}
					} else if tuple.Relation == "*" {
						if tuple.ObjectNamespace == to.ObjectNamespace && tuple.ObjectName == to.ObjectName {
							return true, nil
						}
						// abstract link
						nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
							SubjectSetObjectNamespace: tuple.ObjectNamespace,
							SubjectSetObjectName:      tuple.ObjectName,
						})
						if err != nil {
							return false, err
						}
						set := utils.NewSet[string]()
						for _, nq := range nextQueries {
							set.Add(nq.Relation)
						}
						for _, rel := range set.ToSlice() {
							q.Push(domain.RelationTuple{
								SubjectSetObjectNamespace: tuple.ObjectNamespace,
								SubjectSetObjectName:      tuple.ObjectName,
								SubjectSetRelation:        rel,
							})
						}
					}
				}

				if tuple.ObjectNamespace == "role" {
					nextQuery := domain.RelationTuple{
						SubjectSetObjectNamespace: "role",
						SubjectSetObjectName:      tuple.ObjectName,
					}
					q.Push(nextQuery)
				}
				// push itself
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

func (u *RelationUsecase) ReversedSearch() error {
	return nil
}
