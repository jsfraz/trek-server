package utils

import (
	"gorm.io/gorm"
)

// https://blog.devgenius.io/singleton-pattern-in-go-4faea607ad0f
type Singleton struct {
	Config     *Config
	PostgresDb gorm.DB
}

var instance *Singleton

// Returns singleton instance.
//
//	@return *Singleton
func GetSingleton() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}
