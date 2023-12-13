package usecase

import (
	sqldomain "rbac/domain/infra/sql"
)

type RoleUsecase struct {
	RelationTupleRepo sqldomain.RelationTupleRepository
}

func NewRoleUsecase(relationTupleRepo sqldomain.RelationTupleRepository) *RoleUsecase {
	return &RoleUsecase{
		RelationTupleRepo: relationTupleRepo,
	}
}

func (ru *RoleUsecase) AddPermissionToRole(objnamespace, objname, relation, rolename string) error
func (ru *RoleUsecase) AssignRoleUpRole(childRolename, parentRolename string) error
func (ru *RoleUsecase) ListChildRoles(rolename string) ([]string, error)
func (ru *RoleUsecase) ListRolePermissions(rolename string) ([]string, error)
func (ru *RoleUsecase) ListRoles() ([]string, error)
func (ru *RoleUsecase) GetRoleMembers(rolename string) ([]string, error)
func (ru *RoleUsecase) DeleteRole(rolename string) error
