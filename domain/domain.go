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
	ObjectNamespace string
	ObjectName      string
	Relation        string
}

type Subject struct {
	SubjectNamespace string
	SubjectName      string
	SubjectRelation  string
}
