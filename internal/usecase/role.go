package usecase

import (
	ucdomain "rbac/domain/usecase"
	"rbac/utils"

	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
	zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"

	"gorm.io/gorm"
)

type RoleUsecase struct {
	ZanzibarDagClient   *zclient.ZanzibarDagClient
	RelationUsecaseRepo ucdomain.RelationUsecase
}

func NewRoleUsecase(zanzibarDagClient *zclient.ZanzibarDagClient, relationUsecaseRepo ucdomain.RelationUsecase) *RoleUsecase {
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
	queries := []zanzibardagdom.Relation{
		{
			SubjectNamespace: "role",
			SubjectName:      name,
		},
		{
			ObjectNamespace: "role",
			ObjectName:      name,
		},
	}

	return u.ZanzibarDagClient.DeleteByQueries(queries)
}

func (u *RoleUsecase) AddRelation(objNamespace, objName, relation, rolename string) error {
	return u.ZanzibarDagClient.BatchOperation([]zanzibardagdom.Operation{
		{
			Type: zanzibardagdom.CreateOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  objNamespace,
				ObjectName:       objName,
				Relation:         relation,
				SubjectNamespace: "role",
				SubjectName:      rolename,
				SubjectRelation:  "member",
			},
		},
		{
			Type: zanzibardagdom.CreateOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  objNamespace,
				ObjectName:       objName,
				Relation:         relation,
				SubjectNamespace: "role",
				SubjectName:      rolename,
				SubjectRelation:  "parent",
			},
		},
	})
}

func (u *RoleUsecase) RemoveRelation(objNamespace, objName, relation, rolename string) error {
	return u.ZanzibarDagClient.BatchOperation([]zanzibardagdom.Operation{
		{
			Type: zanzibardagdom.DeleteOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  objNamespace,
				ObjectName:       objName,
				Relation:         relation,
				SubjectNamespace: "role",
				SubjectName:      rolename,
				SubjectRelation:  "member",
			},
		},
		{
			Type: zanzibardagdom.DeleteOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  objNamespace,
				ObjectName:       objName,
				Relation:         relation,
				SubjectNamespace: "role",
				SubjectName:      rolename,
				SubjectRelation:  "parent",
			},
		},
	})
}

// TODO: should batch three operations
func (u *RoleUsecase) AddParent(childRolename, parentRolename string) error {
	if err := u.ZanzibarDagClient.BatchOperation([]zanzibardagdom.Operation{
		{
			Type: zanzibardagdom.CreateOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  "role",
				ObjectName:       childRolename,
				Relation:         "parent",
				SubjectNamespace: "role",
				SubjectName:      parentRolename,
				SubjectRelation:  "member",
			},
		},
		{
			Type: zanzibardagdom.CreateOperation,
			Relation: zanzibardagdom.Relation{
				ObjectNamespace:  "role",
				ObjectName:       childRolename,
				Relation:         "parent",
				SubjectNamespace: "role",
				SubjectName:      parentRolename,
				SubjectRelation:  "parent",
			},
		},
	}); err != nil {
		return err
	}

	relation := zanzibardagdom.Relation{
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

// TODO: if role has no parent, we can delete the modify-permission relation
func (u *RoleUsecase) RemoveParent(childRolename, parentRolename string) error {
	query := zanzibardagdom.Relation{
		ObjectNamespace:  "role",
		ObjectName:       childRolename,
		Relation:         "parent",
		SubjectNamespace: "role",
		SubjectName:      parentRolename,
		SubjectRelation:  "member",
	}

	return u.RelationUsecaseRepo.Delete(query)
}

// TODO: only search for role ns to optimize
func (u *RoleUsecase) GetChildRoles(name string) ([]string, error) {
	relations, err := u.GetAllObjectRelations(name)
	if err != nil {
		return nil, err
	}
	roles := utils.NewSet[string]()
	for _, relation := range relations {
		if relation.ObjectNamespace == "role" {
			roles.Add(relation.ObjectName)
		}
	}

	return roles.ToSlice(), nil
}

func (u *RoleUsecase) GetAllObjectRelations(name string) ([]zanzibardagdom.Relation, error) {
	return u.RelationUsecaseRepo.GetAllObjectRelations(
		zanzibardagdom.Node{
			Namespace: "role",
			Name:      name,
			Relation:  "member",
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

	relations, err := u.RelationUsecaseRepo.Query(query)
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
		SubjectRelation:  "member",
	})
}
