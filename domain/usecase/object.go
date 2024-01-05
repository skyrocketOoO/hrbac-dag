package usecasedomain

type ObjectUsecase interface {
	GetUserRelations(namespace string, name string, relation string) ([]string, error)
	GetRoleRelations(namespace string, name string, relation string) ([]string, error)
}
