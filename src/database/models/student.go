package models

import (
	// persistence layer
	"github.com/jinzhu/gorm"
)

/*
 */
type Student struct {
	gorm.Model

	// data fields
	Name        string `gorm:"unique"`
	Path        string `gorm:"unique"`
	IsReference bool
	Results     []*TaskResult
}

// func (s *Student) BeforeCreate(scope *gorm.Scope) error {
// 	if s.IsReference {
// 		scope.SetColumn("ID", 0)
// 	}
// 	return nil
// }
