package usecasedomain

import zanzibardagdom "rbac/domain/infra/zanzibar-dag"

type RoleUsecase interface {
	GetAll() ([]string, error)
	Delete(name string) error

	AddRelation(objnamespace, objname, relation, rolename string) error
	RemoveRelation(objnamespace, objname, relation, rolename string) error
	AddParent(childRolename, parentRolename string) error
	RemoveParent(childRolename, parentRolename string) error
	GetChildRoles(rolename string) ([]string, error)
	GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error)
	GetMembers(name string) ([]string, error)
	Check(objnamespace, objname, relation, rolename string) (bool, error)
}
