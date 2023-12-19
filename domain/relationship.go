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

type Object struct {
	ObjectNamespace string
	ObjectName      string
	Relation        string
}

type Subject struct {
	SubjectNamespace    string
	SubjectName         string
	SubjectSetNamespace string
	SubjectSetName      string
	SubjectSetRelation  string
}
