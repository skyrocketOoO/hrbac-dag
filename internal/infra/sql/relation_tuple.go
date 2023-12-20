package sql

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"

	"gorm.io/gorm"
)

type RelationTupleRepository struct {
	DB *gorm.DB
}

func NewRelationTupleRepository(db *gorm.DB) *RelationTupleRepository {
	return &RelationTupleRepository{DB: db}
}

func (r *RelationTupleRepository) CreateTuple(tuple domain.RelationTuple) error {
	return r.DB.Create(&tuple).Error
}

func (r *RelationTupleRepository) DeleteTuple(id uint) error {
	return r.DB.Delete(&sqldomain.RelationTuple{}, id).Error
}

func (r *RelationTupleRepository) GetTuples() ([]sqldomain.RelationTuple, error) {
	var tuples []sqldomain.RelationTuple
	if err := r.DB.Find(&tuples).Error; err != nil {
		return nil, err
	}
	return tuples, nil
}

func (r *RelationTupleRepository) QueryTuples(filter domain.RelationTuple) ([]sqldomain.RelationTuple, error) {
	var tuples []sqldomain.RelationTuple
	if err := r.DB.Where(&filter).Find(&tuples).Error; err != nil {
		return nil, err
	}
	return tuples, nil
}

func (r *RelationTupleRepository) GetNamespaces() ([]string, error) {
	var namespaces []string
	if err := r.DB.Model(&sqldomain.RelationTuple{}).Pluck("DISTINCT obj_ns", &namespaces).Error; err != nil {
		return nil, err
	}
	return namespaces, nil
}

func (r *RelationTupleRepository) QueryExactMatchTuples(tuple domain.RelationTuple) ([]sqldomain.RelationTuple, error) {
	var matchingTuples []sqldomain.RelationTuple
	if err := r.DB.Where(&tuple).Find(&matchingTuples).Error; err != nil {
		return nil, err
	}

	return matchingTuples, nil
}

func (r *RelationTupleRepository) DeleteAllTuples() error {
	if err := r.DB.Delete(&sqldomain.RelationTuple{}).Error; err != nil {
		return err
	}
	return nil
}