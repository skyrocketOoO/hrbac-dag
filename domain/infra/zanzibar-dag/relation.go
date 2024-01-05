package zanzibardagdom

type Relation struct {
	ObjectNamespace  string `json:"object_namespace"`
	ObjectName       string `json:"object_name"`
	Relation         string `json:"relation"`
	SubjectNamespace string `json:"subject_namespace"`
	SubjectName      string `json:"subject_name"`
	SubjectRelation  string `json:"subject_relation"`
}

type Node struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Relation  string `json:"relation"`
}

type ZanzibarDagRepository interface {
	GetAll() ([]Relation, error)
	Query(relation Relation) ([]Relation, error)
	Create(relation Relation) error
	Delete(relation Relation) error

	GetAllNamespaces() ([]string, error)
	Check(from Node, to Node) (bool, error)
	GetShortestPath(from Node, to Node) ([]Relation, error)
	GetAllPaths(from Node, to Node) ([][]Relation, error)
	GetAllObjectRelations(subject Node) ([]Relation, error)
	GetAllSubjectRelations(object Node) ([]Relation, error)

	ClearAllRelations() error
}
