package usecasedomain

import zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"

type RoleUsecase interface {
	GetAll() ([]string, error)
	Delete(name string) error

	AddRelation(objnamespace, objname, relation, rolename string) error
	RemoveRelation(objnamespace, objname, relation, rolename string) error
	AddParent(childRolename, parentRolename string) error
	RemoveParent(childRolename, parentRolename string) error
	GetChildRoles(name string) ([]string, error)
	GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error)
	GetMembers(name string) ([]string, error)
	Check(objnamespace, objname, relation, rolename string) (bool, error)
}
