package usecase

import (
	"errors"
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
	ucdomain "rbac/domain/usecase"
	"rbac/utils"
)

type UserUsecase struct {
	RelationTupleRepo   sqldomain.RelationTupleRepository
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewUserUsecase(relationTupleRepo sqldomain.RelationTupleRepository, relationUsecaseRepo ucdomain.RelationUsecase) *UserUsecase {
	return &UserUsecase{
		RelationTupleRepo:   relationTupleRepo,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

func (u *UserUsecase) ListUsers() ([]string, error) {
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("user", "")
	if err != nil {
		return nil, err
	}

	users := utils.NewSet[string]()
	for _, tuple := range tuples {
		if tuple.ObjectNamespace == "user" {
			users.Add(tuple.ObjectName)
		}
		if tuple.SubjectNamespace == "user" {
			users.Add(tuple.SubjectName)
		}
		if tuple.SubjectSetObjectNamespace == "user" {
			users.Add(tuple.SubjectSetObjectName)
		}
	}

	return users.ToSlice(), nil
}

// TODO: this method will check existence after list all relation tuples, but we can optimize to first find
func (u *UserUsecase) GetUser(name string) (string, error) {
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("user", name)
	if err != nil {
		return "", err
	}

	if len(tuples) > 0 {
		return name, nil
	}
	return "", nil
}

func (u *UserUsecase) DeleteUser(name string) error {
	tuples, err := u.RelationUsecaseRepo.ListRelationTuples("user", name)
	if err != nil {
		return err
	}

	for _, tuple := range tuples {
		if err := u.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}
	return nil
}

func (u *UserUsecase) AddRole(username, rolename string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  "role",
		ObjectName:       rolename,
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationTupleRepo.CreateTuple(tuple)
}

func (u *UserUsecase) RemoveRole(username, rolename string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  "role",
		ObjectName:       rolename,
		Relation:         "member",
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	matchedTuples, err := u.RelationTupleRepo.QueryExactMatchTuples(tuple)
	if err != nil {
		return err
	}
	if len(matchedTuples) == 0 {
		return errors.New("the matched tuples is 0")
	}
	for _, tuple := range matchedTuples {
		if err := u.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}

	return nil
}

func (u *UserUsecase) ListRelations(username string) ([]string, error) {
	relations := utils.NewSet[string]()

	firstQuery := domain.RelationTuple{
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	q := utils.NewQueue[domain.RelationTuple]()
	q.Push(firstQuery)
	for !q.IsEmpty() {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			query, err := q.Pop()
			if err != nil {
				return nil, err
			}

			tuples, err := u.RelationTupleRepo.QueryTuples(query)
			if err != nil {
				return nil, err
			}

			for _, tuple := range tuples {
				relations.Add(query.ObjectNamespace + ":" + query.ObjectName + "#" + query.Relation)
				if tuple.ObjectNamespace == "role" {
					// use role to search
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
	return relations.ToSlice(), nil
}

func (u *UserUsecase) AddRelation(username, relation, objectnamespace, objectname string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	return u.RelationUsecaseRepo.Create(tuple)
}

func (u *UserUsecase) RemoveRelation(username, relation, objectnamespace, objectname string) error {
	tuple := domain.RelationTuple{
		ObjectNamespace:  objectnamespace,
		ObjectName:       objectname,
		Relation:         relation,
		SubjectNamespace: "user",
		SubjectName:      username,
	}

	tuples, err := u.RelationTupleRepo.QueryExactMatchTuples(tuple)
	if err != nil {
		return err
	}

	for _, tuple := range tuples {
		if err := u.RelationTupleRepo.DeleteTuple(tuple.ID); err != nil {
			return err
		}
	}
	return nil
}

func (u *UserUsecase) Check(userName, relation, objectNamespace, objectName string) (ok bool, err error) {
	firstQuery := domain.RelationTuple{
		SubjectNamespace: "user",
		SubjectName:      userName,
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
				if tuple.ObjectName == objectName && tuple.ObjectNamespace == objectNamespace && tuple.Relation == relation {
					return true, nil
				}

				if tuple.ObjectNamespace == "role" {
					// use role to search
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
	return false, nil
}
