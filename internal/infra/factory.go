package infra

import (
	"fmt"
	"os"

	zclient "github.com/skyrocketOoO/zanazibar-dag/client"
)

type InfraRepository struct {
	ZanzibarDagClient *zclient.ZanzibarDagClient
}

func NewInfraRepository() (*InfraRepository, error) {
	zanzibarDagUrl := fmt.Sprintf("http://%s:%s", os.Getenv("ZANZIBAR_DAG_HOST"), os.Getenv("ZANZIBAR_DAG_PORT"))
	client, err := zclient.NewZanzibarDagClient(zanzibarDagUrl)
	if err != nil {
		return nil, err
	}
	return &InfraRepository{
		ZanzibarDagClient: client,
	}, nil
}
