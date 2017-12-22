package models

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	// field types
	"github.com/ms-xy/Tutortool/src/database/datatypes"
	"time"
)

type Testcase struct {
	gorm.Model

	// parent task ref
	TaskID uint `gorm:"index"`

	// data fields
	Name   string
	Number int `gorm:"index"`

	Points       int
	LastModified time.Time

	Input      []byte
	Parameters datatypes.Strings `gorm:"type:string"`
	Timeout    time.Duration
	StdoutSize int
	StderrSize int
}
