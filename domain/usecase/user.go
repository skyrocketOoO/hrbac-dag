package usecasedomain

import zanzibardagdom "rbac/domain/infra/zanzibar-dag"

type UserUsecase interface {
	GetAll() ([]string, error)
	Delete(name string) error

	GetRoles(name string) ([]string, error)
	AddRole(username, rolename string) error
	RemoveRole(username, rolename string) error
	GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error)
	AddRelation(username, relation, objnamespace, objname string) error
	RemoveRelation(username, relation, objnamespace, objname string) error
	Check(username, relation, objnamespace, objname string) (bool, error)
}
