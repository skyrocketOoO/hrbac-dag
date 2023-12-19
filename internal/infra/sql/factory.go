package sql

import (
	sqldomain "rbac/domain/infra/sql"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {
	return gorm.Open(
		sqlite.Open("gorm.db"), &gorm.Config{
			Logger: nil,
		},
		// postgres.Open(dsn), &gorm.Config{
		// logger all sql
		//Logger: logger.Default.LogMode(logger.Info),
		// }
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
