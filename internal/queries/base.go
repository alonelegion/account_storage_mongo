package queries

import (
	"gorm.io/gorm"
	"sync"
)

type Query interface {
	GetDB() *gorm.DB
}

type BaseQuery struct {
	db *gorm.DB
}

// Singleton Application implementation
var (
	once sync.Once
	// Current Application
	instance *BaseQuery
)

func InitQuery(db *gorm.DB) *BaseQuery {
	once.Do(func() {
		instance = &BaseQuery{db: db}
	})

	return instance
}

func GetQuery() *BaseQuery {
	return instance
}

func (bq *BaseQuery) GetDB() *gorm.DB {
	return bq.db
}
