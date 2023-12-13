package usecasedomain

type RoleUsecase interface {
	AddPermissionToRole()
	AddRoleToRole()
	GetChildernRole()
	ListRolePermissions()
	ListRoles()
	GetRoleMember()
	DeleteRole()
}
