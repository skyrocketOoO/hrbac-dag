package infra

type InfraRepository struct {
	ZanzibarDagRepo *ZanzibarDagRepository
}

func NewInfraRepository() *InfraRepository {
	return &InfraRepository{
		ZanzibarDagRepo: NewZanzibarDagRepository(),
	}
}
