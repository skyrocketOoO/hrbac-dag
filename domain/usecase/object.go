package usecasedomain

type ObjectUsecase interface {
	ListUserHasRelationOnObject(namespace string, name string, relation string) ([]string, error)
	ListRoleHasWhatRelationOnObject(namespace string, name string, relation string) ([]string, error)
	ListUserOrRoleHasRelationOnObject(namespace string, name string, relation string) ([]string, []string, error)
	ListRelations(namespace, name string) ([]string, error)
}
