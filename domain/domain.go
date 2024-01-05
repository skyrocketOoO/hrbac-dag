package domain

import zanzibardagdom "rbac/domain/infra/zanzibar-dag"

type ErrResponse struct {
	Error string `json:"error"`
}

type StringsResponse struct {
	Data []string `json:"data"`
}

type RelationsResponse struct {
	Data []zanzibardagdom.Relation `json:"data"`
}
