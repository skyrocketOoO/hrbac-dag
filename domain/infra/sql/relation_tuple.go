package sqldomain

import (
	"rbac/domain"
)

type RelationTuple struct {
	ID               uint   `gorm:"primaryKey"`
	ObjectNamespace  string `gorm:"index:idx_object"`
	ObjectName       string `gorm:"index:idx_object"`
	Relation         string `gorm:"index:idx_object"`
	SubjectNamespace string `gorm:"index:idx_subject"`
	SubjectName      string `gorm:"index:idx_subject"`
	SubjectRelation  string `gorm:"index:idx_subject"`
	AllColumns       string `gorm:"uniqueIndex:idx_all_columns"`
}

type RelationTupleRepository interface {
	CreateTuple(tuple domain.RelationTuple) error
	DeleteTuple(id uint) error
	GetAllTuples() ([]RelationTuple, error)
	QueryExactMatchTuples(tuple domain.RelationTuple) ([]RelationTuple, error)
	QueryTuples(query domain.RelationTuple) ([]RelationTuple, error)
	GetNamespaces() ([]string, error)
	DeleteAllTuples() error
}
