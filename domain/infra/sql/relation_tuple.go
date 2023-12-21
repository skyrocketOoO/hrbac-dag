package sqldomain

import (
	"rbac/domain"
)

type RelationTuple struct {
	ID               uint   `gorm:"primaryKey"`
	ObjectNamespace  string `gorm:"uniqueIndex:tuple"`
	ObjectName       string `gorm:"uniqueIndex:tuple"`
	Relation         string `gorm:"uniqueIndex:tuple"`
	SubjectNamespace string `gorm:"uniqueIndex:tuple"`
	SubjectName      string `gorm:"uniqueIndex:tuple"`
	SubjectRelation  string `gorm:"uniqueIndex:tuple"`
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
