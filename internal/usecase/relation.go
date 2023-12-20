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

func (u *RelationUsecase) GetAllRelations() ([]string, error) {
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
	// if tuple exist, return
	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(relationTuple)
	if err != nil {
		return err
	}
	if len(tuples) > 0 {
		return nil
	}

	if err := u.RelationTupleRepo.CreateTuple(relationTuple); err != nil {
		return err
	}
	ok, err := u.hasCycle()
	if err != nil {
		return err
	}
	if ok {
		if err := u.Delete(relationTuple); err != nil {
			return err
		}
		return errors.New("create cycle detected")
	}
	return nil
}

func (u *RelationUsecase) Delete(relationTuple domain.RelationTuple) error {
	matchTuples, err := u.RelationTupleRepo.QueryExactMatchTuples(relationTuple)
	if err != nil {
		return err
	}
	if len(matchTuples) > 1 {
		return errors.New("match tuples > 1")
	}
	return u.RelationTupleRepo.DeleteTuple(matchTuples[0].ID)
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

	return u.Create(tuple)
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

func (u *RelationUsecase) GetAllPaths(relationTuple domain.RelationTuple) ([]string, error) {
	return nil, errors.New("not implemented")
}

func (u *RelationUsecase) GetShortestPath(relationTuple domain.RelationTuple) ([]string, error) {
	return nil, errors.New("not implemented")
}

// use to get all relations based on given attr
func (u *RelationUsecase) QueryExistedRelationTuples(namespace, name string) ([]sqldomain.RelationTuple, error) {
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
	if from.SubjectNamespace == "role" && from.SubjectName == "admin" {
		return true, nil
	}

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

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, _ := q.Pop()
			// fmt.Println("========================query============================")
			// fmt.Printf("%+v\n", query)
			// fmt.Println("=========================================================")
			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				// fmt.Printf("%+v\n", tuple)
				// fmt.Printf("%s : %s # %s\n", tuple.ObjectNamespace, tuple.ObjectName, tuple.Relation)
				if tuple.ObjectNamespace == to.ObjectNamespace {
					if tuple.ObjectName == "*" && tuple.Relation == "*" {
						return true, nil
					} else if tuple.ObjectName == "*" && tuple.Relation != "*" {
						if tuple.Relation == to.Relation {
							return true, nil
						}
					} else if tuple.ObjectName != "*" && tuple.Relation == "*" {
						if tuple.ObjectName == to.ObjectName {
							return true, nil
						}
					}
				}
				if tuple.ObjectNamespace == to.ObjectNamespace && tuple.ObjectName == to.ObjectName && tuple.Relation == to.Relation {
					return true, nil
				}
				// WARNING: This method is used for distinct namespace link scenario
				// object case
				// TODO: the <ns> <*> <*> query will direct find the next layer object, so can directly jump to next
				// if tuple.ObjectNamespace != "role" && tuple.ObjectNamespace != "user" {
				// 	if tuple.ObjectName == "*" && tuple.Relation == "*" {
				// 		if tuple.ObjectNamespace == to.ObjectNamespace {
				// 			return true, nil
				// 		}
				// 		// add abstract link
				// 		nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 			SubjectSetObjectName:      "*",
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set := utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.SubjectSetRelation)
				// 		}
				// 		for _, rel := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 				SubjectSetObjectName:      "*",
				// 				SubjectSetRelation:        rel,
				// 			})
				// 		}

				// 		nextQueries, err = u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 			SubjectSetRelation:        "*",
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set = utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.SubjectSetObjectName)
				// 		}
				// 		for _, name := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 				SubjectSetObjectName:      name,
				// 				SubjectSetRelation:        "*",
				// 			})
				// 		}
				// 	} else if tuple.ObjectName == "*" {
				// 		if tuple.ObjectNamespace == to.ObjectNamespace && tuple.Relation == to.Relation {
				// 			return true, nil
				// 		}
				// 		// abstract link
				// 		nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 			SubjectSetRelation:        tuple.Relation,
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set := utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.ObjectName)
				// 		}
				// 		for _, name := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 				SubjectSetObjectName:      name,
				// 				SubjectSetRelation:        tuple.Relation,
				// 			})
				// 		}
				// 	} else if tuple.Relation == "*" {
				// 		if tuple.ObjectNamespace == to.ObjectNamespace && tuple.ObjectName == to.ObjectName {
				// 			return true, nil
				// 		}
				// 		// abstract link
				// 		nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 			SubjectSetObjectName:      tuple.ObjectName,
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set := utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.Relation)
				// 		}
				// 		for _, rel := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectSetObjectNamespace: tuple.ObjectNamespace,
				// 				SubjectSetObjectName:      tuple.ObjectName,
				// 				SubjectSetRelation:        rel,
				// 			})
				// 		}
				// 	}
				// }

				if tuple.ObjectNamespace == "role" {
					nextQuery := domain.RelationTuple{
						SubjectNamespace: "role",
						SubjectName:      tuple.ObjectName,
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
	return errors.New("not implemented")
}

func (u *RelationUsecase) detectCycle(node domain.Object, visited *utils.Set[domain.Object], recursionStack *utils.Set[domain.Object]) (bool, error) {
	visited.Add(node)
	recursionStack.Add(node)

	query := domain.RelationTuple{}
	if node.Relation != "" {
		query.SubjectSetObjectNamespace = node.ObjectNamespace
		query.SubjectSetObjectName = node.ObjectName
		query.SubjectSetRelation = node.Relation
	} else {
		query.SubjectNamespace = node.ObjectNamespace
		query.SubjectName = node.ObjectName
	}
	neighbors, err := u.RelationTupleRepo.QueryTuples(query)
	if err != nil {
		return false, err
	}
	for _, neighbor := range neighbors {
		var object domain.Object
		if neighbor.ObjectNamespace != "role" && neighbor.ObjectNamespace != "user" {
			object.ObjectNamespace = neighbor.ObjectNamespace
			object.ObjectName = neighbor.ObjectName
			object.Relation = neighbor.Relation
		} else {
			object.ObjectNamespace = neighbor.ObjectNamespace
			object.ObjectName = neighbor.ObjectName
		}
		if !visited.Exist(object) {
			ok, err := u.detectCycle(object, visited, recursionStack)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		} else if recursionStack.Exist(object) {
			return true, nil
		}
	}

	recursionStack.Remove(node)
	return false, nil
}

func (u *RelationUsecase) hasCycle() (bool, error) {
	visited := utils.NewSet[domain.Object]()
	recursionStack := utils.NewSet[domain.Object]()

	allTuples, err := u.RelationTupleRepo.GetTuples()
	if err != nil {
		return false, err
	}

	for _, tuple := range allTuples {
		var object domain.Object
		if tuple.SubjectNamespace != "" {
			object.ObjectNamespace = tuple.SubjectNamespace
			object.ObjectName = tuple.SubjectName
		} else {
			object.ObjectNamespace = tuple.SubjectSetObjectNamespace
			object.ObjectName = tuple.SubjectSetObjectName
			object.Relation = tuple.SubjectSetRelation
		}

		if !visited.Exist(object) {
			ok, err := u.detectCycle(object, &visited, &recursionStack)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
	}
	return false, nil
}

func (u *RelationUsecase) FindAllObjectRelations(from domain.Subject) ([]string, error) {
	if from.SubjectNamespace == "role" && from.SubjectName == "admin" {
		return nil, nil
	}
	objectRelations := utils.NewSet[string]()

	var firstQuery domain.RelationTuple
	if from.SubjectNamespace != "" {
		firstQuery.SubjectNamespace = from.SubjectNamespace
		firstQuery.SubjectName = from.SubjectName
	} else if from.SubjectSetNamespace != "" {
		firstQuery.SubjectSetObjectNamespace = from.SubjectSetNamespace
		firstQuery.SubjectSetObjectName = from.SubjectSetName
		firstQuery.SubjectSetRelation = from.SubjectSetRelation
	} else {
		return nil, errors.New("subject error")
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, _ := q.Pop()
			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, err
			}

			for _, tuple := range tuples {
				if tuple.ObjectNamespace != "role" && tuple.ObjectNamespace != "user" {
					relationStr := utils.RelationTupleToString(utils.ConvertRelationTuple(sqldomain.RelationTuple{
						ObjectNamespace: tuple.ObjectNamespace,
						ObjectName:      tuple.ObjectName,
						Relation:        tuple.Relation,
					}))
					objectRelations.Add(relationStr)
				}
				if tuple.ObjectNamespace == "role" {
					nextQuery := domain.RelationTuple{
						SubjectNamespace: "role",
						SubjectName:      tuple.ObjectName,
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

	return objectRelations.ToSlice(), nil
}
