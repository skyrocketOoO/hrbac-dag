package sqldomain

import "gorm.io/gorm"

type RelationTuple struct {
	gorm.Model
	ObjNS          string `gorm:"uniqueIndex:tuple"`
	ObjName        string `gorm:"uniqueIndex:tuple"`
	Relation       string `gorm:"uniqueIndex:tuple"`
	SubNS          string `gorm:"uniqueIndex:tuple"`
	SubName        string `gorm:"uniqueIndex:tuple"`
	SubSetObjNS    string `gorm:"uniqueIndex:tuple"`
	SubSetObjName  string `gorm:"uniqueIndex:tuple"`
	SubSetRelation string `gorm:"uniqueIndex:tuple"`
}

type RelationTupleRepository interface {
	CreateTuple(tuple RelationTuple) error
	DeleteTuple(id uint) error
	GetTuples() ([]RelationTuple, error)
	QueryTuples(filter RelationTuple) ([]RelationTuple, error)
	GetNamespaces() ([]string, error)
}
