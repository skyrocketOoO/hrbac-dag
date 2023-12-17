package domain

type RelationTuple struct {
	ObjectNamespace           string
	ObjectName                string
	Relation                  string
	SubjectNamespace          string
	SubjectName               string
	SubjectSetObjectNamespace string
	SubjectSetObjectName      string
	SubjectSetRelation        string
}
