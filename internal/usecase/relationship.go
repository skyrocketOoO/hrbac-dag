package usecase

import (
	sqldomain "rbac/domain/infra/sql"
)

type RelationshipUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRelationshipUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RelationshipUsecase {
	return &RelationshipUsecase{RelationTupleRepo: relationTupleRepo}
}

func (ru *RelationshipUsecase) CreateTuple(relationTuple sqldomain.RelationTuple) error {
	return ru.RelationTupleRepo.CreateTuple(relationTuple)
}

func (ru *RelationshipUsecase) DeleteTuple(id uint) error {
	return ru.RelationTupleRepo.DeleteTuple(id)
}

func (ru *RelationshipUsecase) GetTuples() ([]sqldomain.RelationTuple, error) {
	return ru.RelationTupleRepo.GetTuples()
}

func (ru *RelationshipUsecase) QueryTuples(filter sqldomain.RelationTuple) ([]sqldomain.RelationTuple, error) {
	return ru.RelationTupleRepo.QueryTuples(filter)
}

func (ru *RelationshipUsecase) GetNamespaces() ([]string, error) {
	return ru.RelationTupleRepo.GetNamespaces()
}
