package usecasedomain

type UserUsecase interface {
	ListUsers() ([]string, error)
	GetUser(name string) (string, error)
	DeleteUser(name string) error

	AddRole(username, rolename string) error
	RemoveRole(username, rolename string) error
	ListRelations(name string) ([]string, error)
	AddRelation(username, relation, objectnamespace, objectname string) error
	RemoveRelation(username, relation, objectnamespace, objectname string) error
	Check(username, relation, objectnamespace, objectname string) (bool, error)
}
