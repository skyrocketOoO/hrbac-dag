package infra

import (
	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
)

type ZanzibarDagRepository struct {
	Url string
	zclient.ZanzibarDagClient
}
