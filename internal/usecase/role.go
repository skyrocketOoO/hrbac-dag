package usecase

import (
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
	ucdomain "rbac/domain/usecase"
	"rbac/utils"

	"gorm.io/gorm"
)

type RoleUsecase struct {
	ZanzibarDagClient   zanzibardagdom.ZanzibarDagRepository
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewRoleUsecase(zanzibarDagClient zanzibardagdom.ZanzibarDagRepository, relationUsecaseRepo ucdomain.RelationUsecase) *RoleUsecase {
	return &RoleUsecase{
		ZanzibarDagClient:   zanzibarDagClient,
		RelationUsecaseRepo: relationUsecaseRepo,
	}
}

// TODO: seems can optimize
func (u *RoleUsecase) GetAll() ([]string, error) {
	relations, err := u.RelationUsecaseRepo.QueryExistedRelations("role", "")
	if err != nil {
		return nil, err
	}

	roles := utils.NewSet[string]()
	for _, relation := range relations {
		if relation.ObjectNamespace == "role" {
			roles.Add(relation.ObjectName)
		}
		if relation.SubjectNamespace == "role" {
			roles.Add(relation.SubjectName)
		}
		if relation.SubjectNamespace == "role" {
			roles.Add(relation.SubjectName)
		}
	}

	return roles.ToSlice(), nil
}

func (u *RoleUsecase) Delete(name string) error {
	relations, err := u.RelationUsecaseRepo.QueryExistedRelations("role", name)
	if err != nil {
		return err
	}

	for _, relation := range relations {
		if err := u.ZanzibarDagClient.Delete(relation); err != nil {
			return err
		}
	}

	return nil
}

func (u *RoleUsecase) AddRelation(objNamespace, objName, relation, rolename string) error {
	rel := zanzibardagdom.Relation{
		ObjectNamespace:  objNamespace,
		ObjectName:       objName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      rolename,
	}

	return u.RelationUsecaseRepo.Create(rel)
}

func (u *RoleUsecase) RemoveRelation(objnamespace, objectName, relation, rolename string) error {
	rel := zanzibardagdom.Relation{
		ObjectNamespace:  objnamespace,
		ObjectName:       objectName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      rolename,
	}

	return u.RelationUsecaseRepo.Delete(rel)
}

func (u *RoleUsecase) AddParent(childRolename, parentRolename string) error {
	relation := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
		SubjectRelation:  "member",
	}
	if err := u.RelationUsecaseRepo.Create(relation); err != nil {
		return err
	}

	relation = zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
		SubjectRelation:  "parent",
	}
	if err := u.RelationUsecaseRepo.Create(relation); err != nil {
		return err
	}

	relation = zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "modify-permission",
		SubjectNamespace: "role",
		SubjectName:      childRolename,
		SubjectRelation:  "parent",
	}

	err := u.RelationUsecaseRepo.Create(relation)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return nil
		}
		return err
	}
	return nil
}

func (u *RoleUsecase) RemoveParent(childRolename, parentRolename string) error {
	query := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
	}

	return u.RelationUsecaseRepo.Delete(query)
}

func (u *RoleUsecase) GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error) {
	return u.RelationUsecaseRepo.GetAllObjectRelations(
		zanzibardagdom.Node{
			Namespace: "role",
			Name:      name,
		},
	)
}

func (u *RoleUsecase) GetMembers(name string) ([]string, error) {
	query := zanzibardagdom.Relation{
		ObjectNamespace: "role",
		ObjectName:      name,
		Relation:        "member",
	}

	users := utils.NewSet[string]()

	relations, err := u.RelationTupleRepo.QueryTuples(query)
	if err != nil {
		return nil, err
	}
	for _, relation := range relations {
		if relation.SubjectNamespace == "user" {
			users.Add(relation.SubjectName)
		}
	}
	return users.ToSlice(), nil
}

func (u *RoleUsecase) Check(objectNamespace, objectName, relation, roleName string) (bool, error) {
	return u.RelationUsecaseRepo.Check(zanzibardagdom.Relation{
		ObjectNamespace:  objectNamespace,
		ObjectName:       objectName,
		Relation:         relation,
		SubjectNamespace: "role",
		SubjectName:      roleName,
	})
}
