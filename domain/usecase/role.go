package usecasedomain

type RoleUsecase interface {
	AddPermissionToRole(objnamespace, objname, relation, rolename string) error
	AssignRoleUpRole(childRolename, parentRolename string) error
	ListChildRoles(rolename string) ([]string, error)
	ListRolePermissions(rolename string) ([]string, error)
	ListRoles() ([]string, error)
	GetRoleMembers(rolename string) ([]string, error)
	DeleteRole(rolename string) error
}
