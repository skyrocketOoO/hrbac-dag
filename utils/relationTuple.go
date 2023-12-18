package utils

import (
	"rbac/domain"
	sqldomain "rbac/domain/infra/sql"
)

func RelationTupleToString(tuple domain.RelationTuple) string {
	res := tuple.ObjectNamespace + ":" + tuple.ObjectName + "#" + tuple.Relation + "@"
	if tuple.SubjectNamespace != "" {
		res += tuple.SubjectNamespace + ":" + tuple.SubjectName
	} else {
		res += "(" + tuple.SubjectSetObjectNamespace + ":" + tuple.SubjectSetObjectName
		res += "#" + tuple.SubjectSetRelation + ")"
	}

	return res
}

func ConvertRelationTuple(in sqldomain.RelationTuple) domain.RelationTuple {
	return domain.RelationTuple{
		ObjectNamespace:           in.ObjectNamespace,
		ObjectName:                in.ObjectName,
		Relation:                  in.Relation,
		SubjectNamespace:          in.SubjectNamespace,
		SubjectName:               in.SubjectName,
		SubjectSetObjectNamespace: in.SubjectSetObjectNamespace,
		SubjectSetObjectName:      in.SubjectSetObjectName,
		SubjectSetRelation:        in.SubjectSetRelation,
	}
}
