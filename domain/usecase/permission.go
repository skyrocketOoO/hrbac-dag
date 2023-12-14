package usecasedomain

type PermissionUsecase interface {
	CheckUserPermission(objNS, objName, Permission, username string) (bool, error)
}
