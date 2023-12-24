package utils

import (
	"errors"
	"rbac/domain"
)

func HasReserveWord(word string, sort string) bool {
	switch sort {
	case "relation":
		if word == "member" || word == "parent" {
			return true
		}
	case "namespace":
		if word == "role" || word == "user" {
			return true
		}
	case "name":
		if word == "admin" {
			return true
		}
	default:
	}

	return false
}

func CheckReserveWordInTuple(tuple domain.RelationTuple) error {
	if HasReserveWord(tuple.ObjectNamespace, "namespace") ||
		HasReserveWord(tuple.ObjectName, "name") ||
		HasReserveWord(tuple.Relation, "relation") ||
		HasReserveWord(tuple.SubjectNamespace, "namespace") ||
		HasReserveWord(tuple.SubjectName, "name") ||
		HasReserveWord(tuple.SubjectRelation, "relation") {
		return errors.New("has reserved word")
	}
	return nil
}
