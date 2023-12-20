package sqldomain

import (
	"rbac/domain"

	"gorm.io/gorm"
)

type RelationTuple struct {
	gorm.Model
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
	GetTuples() ([]RelationTuple, error)
	QueryExactMatchTuples(tuple domain.RelationTuple) ([]RelationTuple, error)
	QueryTuples(query domain.RelationTuple) ([]RelationTuple, error)
	GetNamespaces() ([]string, error)
	DeleteAllTuples() error
}
