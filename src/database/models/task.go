package models

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	// field types
	"github.com/ms-xy/Tutortool/src/database/datatypes"
	"time"
	// utility
	"sync"
)

type Task struct {
	gorm.Model

	// data fields
	Path         string
	Name         string
	LastModified time.Time

	NamingSchema string
	DueDate      time.Time

	RefImplPath string

	GccParams       datatypes.Strings          `gorm:"type:string"`
	GccFiles        datatypes.Strings          `gorm:"type:string"`
	GccReplacements datatypes.String2StringMap `gorm:"type:string"`
	/*
		replacements is a map of:
		student-submission-relative-filepath -> reference-impl-relative-filepath

		e.g.: to replace the students main.c with the ref-impls main.c, supply:
		main.c -> main.c
	*/

	RunTimeout    time.Duration
	RunStdoutSize int
	RunStderrSize int

	Testcases []*Testcase

	Sources datatypes.Strings `gorm:"type:string"`

	Lock sync.Mutex `gorm:"-"` // must not be persistet
}
