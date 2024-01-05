package utils

import (
	"errors"
	zanzibardagdom "rbac/domain/infra/zanzibar-dag"
)

func HasReserveWord(word string, sort string) bool {
	switch sort {
	// case "relation":
	// 	if word == "member" || word == "parent" || word == "modify-permission" {
	// 		return true
	// 	}
	// case "namespace":
	// 	if word == "role" || word == "user" {
	// 		return true
	// 	}
	case "name":
		if word == "admin" {
			return true
		}
	default:
	}

	return false
}

func ValidateReserveWord(relation zanzibardagdom.Relation) error {
	if HasReserveWord(relation.ObjectNamespace, "namespace") ||
		HasReserveWord(relation.ObjectName, "name") ||
		HasReserveWord(relation.Relation, "relation") ||
		HasReserveWord(relation.SubjectNamespace, "namespace") ||
		HasReserveWord(relation.SubjectName, "name") ||
		HasReserveWord(relation.SubjectRelation, "relation") {
		return errors.New("has reserved word")
	}
	return nil
}
