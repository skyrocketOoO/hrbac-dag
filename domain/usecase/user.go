package usecasedomain

type UserUsecase interface {
	ListUsers() ([]string, error)
	AddUserToRole(username, rolename string) error
	RemoveUserFromRole(username, rolename string)
	ListUserPermissions(username string) ([]string, error)
	AddPermissionToUser(username, relation, objectnamespace, objectname string) error
}
