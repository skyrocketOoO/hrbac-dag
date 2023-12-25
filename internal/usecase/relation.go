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
	tuples, err := u.RelationTupleRepo.GetAllTuples()
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
	if err := utils.CheckReserveWordInTuple(relationTuple); err != nil {
		return err
	}

	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(relationTuple)
	if err != nil {
		return err
	}
	if len(tuples) > 0 {
		return errors.New("tuple exist")
	}

	ok, err := u.Check(domain.RelationTuple{
		ObjectNamespace:  relationTuple.SubjectNamespace,
		ObjectName:       relationTuple.SubjectName,
		Relation:         relationTuple.SubjectRelation,
		SubjectNamespace: relationTuple.ObjectNamespace,
		SubjectName:      relationTuple.ObjectName,
		SubjectRelation:  relationTuple.Relation,
	})
	if err != nil {
		return err
	}
	if ok {
		return errors.New("create cycle detected")
	}
	return u.RelationTupleRepo.CreateTuple(relationTuple)
}

func (u *RelationUsecase) Delete(relationTuple domain.RelationTuple) error {
	if err := utils.CheckReserveWordInTuple(relationTuple); err != nil {
		return err
	}

	matchTuples, err := u.RelationTupleRepo.QueryExactMatchTuples(relationTuple)
	if err != nil {
		return err
	}
	if len(matchTuples) > 1 {
		return errors.New("match tuples > 1")
	} else if len(matchTuples) == 0 {
		return errors.New("relation not found")
	}
	return u.RelationTupleRepo.DeleteTuple(matchTuples[0].ID)
}

func (u *RelationUsecase) AddLink(tuple domain.RelationTuple) error {
	return u.Create(tuple)
}

func (u *RelationUsecase) RemoveLink(tuple domain.RelationTuple) error {
	return u.Delete(tuple)
}

func (u *RelationUsecase) Check(relationTuple domain.RelationTuple) (bool, error) {
	from := domain.Subject{
		Namespace: relationTuple.SubjectNamespace,
		Name:      relationTuple.SubjectName,
		Relation:  relationTuple.SubjectRelation,
	}
	to := domain.Object{
		Namespace: relationTuple.ObjectNamespace,
		Name:      relationTuple.ObjectName,
		Relation:  relationTuple.Relation,
	}
	if from.Namespace == "role" && from.Name == "admin" {
		return true, nil
	}

	visited := utils.NewSet[domain.RelationTuple]()

	firstQuery := domain.RelationTuple{
		SubjectNamespace: from.Namespace,
		SubjectName:      from.Name,
		SubjectRelation:  from.Relation,
	}
	q := utils.NewQueue[domain.RelationTuple]()
	visited.Add(firstQuery)
	q.Push(firstQuery)
	if firstQuery.SubjectNamespace == "role" {
		firstQuery.SubjectRelation = ""
		q.Push(firstQuery)
		visited.Add(firstQuery)
	}
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, _ := q.Pop()
			// fmt.Println("========================query============================")
			// fmt.Printf("%+v\n", query)
			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return false, err
			}

			for _, tuple := range tuples {
				// fmt.Println("===========================tuple=========================")
				// fmt.Printf("%+v\n", tuple)
				// fmt.Printf("%s : %s # %s\n", tuple.ObjectNamespace, tuple.ObjectName, tuple.Relation)
				if tuple.ObjectNamespace == to.Namespace {
					if tuple.ObjectName == "*" && tuple.Relation == "*" {
						return true, nil
					} else if tuple.ObjectName == "*" && tuple.Relation != "*" {
						if tuple.Relation == to.Relation {
							return true, nil
						}
					} else if tuple.ObjectName != "*" && tuple.Relation == "*" {
						if tuple.ObjectName == to.Name {
							return true, nil
						}
					}
				}
				if tuple.ObjectNamespace == to.Namespace && tuple.ObjectName == to.Name && tuple.Relation == to.Relation {
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
				// 			SubjectNamespace: tuple.ObjectNamespace,
				// 			SubjectName:      "*",
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set := utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.SubjectRelation)
				// 		}
				// 		for _, rel := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectNamespace: tuple.ObjectNamespace,
				// 				SubjectName:      "*",
				// 				SubjectRelation:        rel,
				// 			})
				// 		}

				// 		nextQueries, err = u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectNamespace: tuple.ObjectNamespace,
				// 			SubjectRelation:        "*",
				// 		})
				// 		if err != nil {
				// 			return false, err
				// 		}
				// 		set = utils.NewSet[string]()
				// 		for _, nq := range nextQueries {
				// 			set.Add(nq.SubjectName)
				// 		}
				// 		for _, name := range set.ToSlice() {
				// 			q.Push(domain.RelationTuple{
				// 				SubjectNamespace: tuple.ObjectNamespace,
				// 				SubjectName:      name,
				// 				SubjectRelation:        "*",
				// 			})
				// 		}
				// 	} else if tuple.ObjectName == "*" {
				// 		if tuple.ObjectNamespace == to.ObjectNamespace && tuple.Relation == to.Relation {
				// 			return true, nil
				// 		}
				// 		// abstract link
				// 		nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectNamespace: tuple.ObjectNamespace,
				// 			SubjectRelation:        tuple.Relation,
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
				// 				SubjectNamespace: tuple.ObjectNamespace,
				// 				SubjectName:      name,
				// 				SubjectRelation:        tuple.Relation,
				// 			})
				// 		}
				// 	} else if tuple.Relation == "*" {
				// 		if tuple.ObjectNamespace == to.ObjectNamespace && tuple.ObjectName == to.ObjectName {
				// 			return true, nil
				// 		}
				// 		// abstract link
				// 		nextQueries, err := u.RelationTupleRepo.QueryTuples(domain.RelationTuple{
				// 			SubjectNamespace: tuple.ObjectNamespace,
				// 			SubjectName:      tuple.ObjectName,
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
				// 				SubjectNamespace: tuple.ObjectNamespace,
				// 				SubjectName:      tuple.ObjectName,
				// 				SubjectRelation:        rel,
				// 			})
				// 		}
				// 	}
				// }

				if tuple.ObjectNamespace == "role" {
					nextQuery := domain.RelationTuple{
						SubjectNamespace: "role",
						SubjectName:      tuple.ObjectName,
					}
					if !visited.Exist(nextQuery) {
						visited.Add(nextQuery)
						q.Push(nextQuery)
					}
				}
				nextQuery := domain.RelationTuple{
					SubjectNamespace: tuple.ObjectNamespace,
					SubjectName:      tuple.ObjectName,
					SubjectRelation:  tuple.Relation,
				}
				if !visited.Exist(nextQuery) {
					visited.Add(nextQuery)
					q.Push(nextQuery)
				}
			}
		}
	}

	return false, nil
}

func (u *RelationUsecase) GetAllPaths(relationTuple domain.RelationTuple) ([]string, error) {
	if err := utils.CheckReserveWordInTuple(relationTuple); err != nil {
		return nil, err
	}
	return nil, errors.New("not implemented")
}

func (u *RelationUsecase) GetShortestPath(relationTuple domain.RelationTuple) ([]string, error) {
	if err := utils.CheckReserveWordInTuple(relationTuple); err != nil {
		return nil, err
	}
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

func (u *RelationUsecase) ReversedSearch() error {
	return errors.New("not implemented")
}

// func (u *RelationUsecase) detectCycle(from domain.Object, visited *utils.Set[domain.Object], recursionStack *utils.Set[domain.Object]) (bool, error) {
// 	visited.Add(from)
// 	recursionStack.Add(from)

// 	query := domain.RelationTuple{
// 		SubjectNamespace: from.Namespace,
// 		SubjectName:      from.Name,
// 		SubjectRelation:  from.Relation,
// 	}
// 	neighbors, err := u.RelationTupleRepo.QueryTuples(query)
// 	if err != nil {
// 		return false, err
// 	}
// 	for _, neighbor := range neighbors {
// 		var object domain.Object
// 		object.Namespace = neighbor.ObjectNamespace
// 		object.Name = neighbor.ObjectName
// 		object.Relation = neighbor.Relation
// 		if !visited.Exist(object) {
// 			ok, err := u.detectCycle(object, visited, recursionStack)
// 			if err != nil {
// 				return false, err
// 			}
// 			if ok {
// 				return true, nil
// 			}
// 		} else if recursionStack.Exist(object) {
// 			return true, nil
// 		}
// 	}

// 	recursionStack.Remove(from)
// 	return false, nil
// }

// // we only need to start from new edge, if A -> B, search from B
// func (u *RelationUsecase) hasCycle() (bool, error) {
// 	visited := utils.NewSet[domain.Object]()
// 	recursionStack := utils.NewSet[domain.Object]()

// 	allTuples, err := u.RelationTupleRepo.GetAllTuples()
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, tuple := range allTuples {
// 		var object domain.Object
// 		object.Namespace = tuple.SubjectNamespace
// 		object.Name = tuple.SubjectName
// 		object.Relation = tuple.SubjectRelation

// 		if !visited.Exist(object) {
// 			ok, err := u.detectCycle(object, &visited, &recursionStack)
// 			if err != nil {
// 				return false, err
// 			}
// 			if ok {
// 				return true, nil
// 			}
// 		}
// 	}
// 	return false, nil
// }

func (u *RelationUsecase) FindAllObjectRelations(from domain.Subject) ([]string, error) {
	if err := utils.CheckReserveWordInTuple(domain.RelationTuple{
		SubjectNamespace: from.Namespace,
		SubjectName:      from.Name,
		SubjectRelation:  from.Relation,
	}); err != nil {
		return nil, err
	}
	if from.Namespace == "role" && from.Name == "admin" {
		return nil, nil
	}
	objectRelations := utils.NewSet[string]()
	visited := utils.NewSet[domain.RelationTuple]()

	firstQuery := domain.RelationTuple{
		SubjectNamespace: from.Namespace,
		SubjectName:      from.Name,
		SubjectRelation:  from.Relation,
	}
	q := utils.NewQueue[domain.RelationTuple]()
	visited.Add(firstQuery)
	q.Push(firstQuery)
	if firstQuery.SubjectNamespace == "role" {
		firstQuery.SubjectRelation = ""
		q.Push(firstQuery)
		visited.Add(firstQuery)
	}
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
					if !visited.Exist(nextQuery) {
						visited.Add(nextQuery)
						q.Push(nextQuery)
					}
				}
				nextQuery := domain.RelationTuple{
					SubjectNamespace: tuple.ObjectNamespace,
					SubjectName:      tuple.ObjectName,
					SubjectRelation:  tuple.Relation,
				}
				if !visited.Exist(nextQuery) {
					visited.Add(nextQuery)
					q.Push(nextQuery)
				}
			}
		}
	}

	return objectRelations.ToSlice(), nil
}
