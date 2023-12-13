package usecasedomain

type ObjectUsecase interface {
	LinkPermission(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error
	ListWhoHasRelationOnObject(namespace string, name string, relation string) ([]string, error)
	ListRolesHasRelationOnObject(namespace string, name string, relation string) ([]string, error)
	ListWhoOrRoleHasRelationOnObject(namespace string, name string, relation string) ([]string, []string, error)
	ListAllPermissions(namespace, name string) ([]string, error)
}
