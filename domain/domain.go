package domain

type RelationTuple struct {
	ObjectNamespace  string
	ObjectName       string
	Relation         string
	SubjectNamespace string
	SubjectName      string
	SubjectRelation  string
}

type Object struct {
	Namespace string
	Name      string
	Relation  string
}

type Subject struct {
	Namespace string
	Name      string
	Relation  string
}
