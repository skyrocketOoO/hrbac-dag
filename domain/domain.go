package domain

import zanzibardagdom "github.com/skyrocketOoO/zanazibar-dag/domain"

type ErrResponse struct {
	Error string `json:"error"`
}

type StringsResponse struct {
	Data []string `json:"data"`
}

type RelationsResponse struct {
	Data []zanzibardagdom.Relation `json:"data"`
}
