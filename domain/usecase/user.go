package usecasedomain

type UserUsecase interface {
	AddUserToRole()
	RemoveUserFromRole()
	ListUserPermissions()
	AddPermissionToUser()
}
