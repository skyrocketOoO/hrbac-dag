package sql

import (
	sqldomain "rbac/domain/infra/sql"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {
	return gorm.Open(
		// sqlite.Open("gorm.db"), &gorm.Config{
		// 	Logger: nil,
		// },
		postgres.Open(viper.GetString("database.postgres.dsn")), &gorm.Config{
			Logger: nil,
		},
	)
}

type OrmRepository struct {
	RelationshipRepo RelationTupleRepository
}

func NewOrmRepository(db *gorm.DB) (*OrmRepository, error) {
	if err := db.AutoMigrate(&sqldomain.RelationTuple{}); err != nil {
		return nil, err
	}

	return &OrmRepository{
		RelationshipRepo: *NewRelationTupleRepository(db),
	}, nil
}
