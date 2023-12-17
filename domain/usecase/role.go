package usecasedomain

type RoleUsecase interface {
	ListRoles() ([]string, error)
	GetRole(name string) (string, error)
	DeleteRole(name string) error

	AddRelation(objnamespace, objname, relation, rolename string) error
	RemoveRelation(objnamespace, objname, relation, rolename string) error
	AddParent(childRolename, parentRolename string) error
	RemoveParent(childRolename, parentRolename string) error
	// ListChildRoles(rolename string) ([]string, error)
	ListRelations(name string) ([]string, error)
	GetMembers(name string) ([]string, error)
	Check(objnamespace, objname, relation, rolename string) (bool, error)
}
