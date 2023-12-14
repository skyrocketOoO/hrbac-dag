package usecasedomain

type ObjectUsecase interface {
	LinkPermission(objnamespace, objname, relation, subjnamespace, subjname, subjrelation string) error
	ListWhoHasPermissionOnObject(namespace string, name string, relation string) ([]string, error)
	ListRolesHasWhatPermissonOnObject(namespace string, name string, relation string) ([]string, error)
	ListWhoOrRoleHasPermissionOnObject(namespace string, name string, relation string) ([]string, []string, error)
	ListAllPermissions(namespace, name string) ([]string, error)
}
