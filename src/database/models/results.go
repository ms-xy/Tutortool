package models

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	// field types
	"github.com/ms-xy/execute"
)

type TaskResult struct {
	gorm.Model
	// ref
	StudentID uint // 0 = reference_impl
	TaskID    uint
	// data fields
	Points     int
	Comment    string `gorm:"size:2000"`
	GccResult  *GccResult
	RunResults []*RunResult
}

type RunResult struct {
	gorm.Model
	// ref
	TaskResultID   uint
	TestcaseNumber int
	TestcaseID     uint
	// data fields
	Points int
	execute.ExecResult
}

type GccResult struct {
	gorm.Model
	// ref
	TaskResultID uint
	// data fields
	execute.ExecResult
}
