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

type SearchCondition struct {
	In Compare `json:"in"`
}

type Compare struct {
	Namespaces []string `json:"namespaces"`
	Names      []string `json:"names"`
	Relations  []string `json:"relations"`
}

type CollectCondition struct {
	In Compare `json:"in"`
}

type Action string

const (
	CreateOperation Action = "create"
	DeleteOperation Action = "delete"
)

type Operation struct {
	Type     Action   `json:"action"`
	Relation Relation `json:"relation"`
}

type ZanzibarDagRepository interface {
	GetAll() ([]Relation, error)
	Query(relation Relation) ([]Relation, error)
	Create(relation Relation) error
	Delete(relation Relation) error
	DeleteByQueries(queries []Relation) error
	BatchOperation(operations []Operation) error

	GetAllNamespaces() ([]string, error)
	Check(from Node, to Node, searchCond SearchCondition) (bool, error)
	GetShortestPath(from Node, to Node, searchCond SearchCondition) ([]Relation, error)
	GetAllPaths(from Node, to Node, searchCond SearchCondition) ([][]Relation, error)
	GetAllObjectRelations(subject Node, searchCond SearchCondition, collectCond CollectCondition) ([]Relation, error)
	GetAllSubjectRelations(object Node, searchCond SearchCondition, collectCond CollectCondition) ([]Relation, error)

	ClearAllRelations() error
}
