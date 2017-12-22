package models

import (
	"github.com/ms-xy/Tutortool/src/database"
)

func init() {
	db := database.GetInstance()

	tables := []interface{}{
		(*Student)(nil),

		(*Task)(nil),

		(*Testcase)(nil),

		(*TaskResult)(nil),
		(*GccResult)(nil),
		(*RunResult)(nil),
	}

	for _, table := range tables {
		db.AutoMigrate(table)
	}
}
