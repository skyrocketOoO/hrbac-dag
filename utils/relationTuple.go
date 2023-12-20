package utils

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
)

func RelationTupleToString(tuple domain.RelationTuple) string {
	res := tuple.ObjectNamespace + ":" + tuple.ObjectName + "#" + tuple.Relation
	res += "@" + tuple.SubjectNamespace + ":" + tuple.SubjectName
	if tuple.SubjectRelation != "" {
		res += "#" + tuple.SubjectRelation
	}

	return res
}

func ConvertRelationTuple(in sqldomain.RelationTuple) domain.RelationTuple {
	return domain.RelationTuple{
		ObjectNamespace:  in.ObjectNamespace,
		ObjectName:       in.ObjectName,
		Relation:         in.Relation,
		SubjectNamespace: in.SubjectNamespace,
		SubjectName:      in.SubjectName,
		SubjectRelation:  in.SubjectRelation,
	}
}
