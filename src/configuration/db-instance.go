package configuration

import (
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"
)

var (
	db *gorm.DB
)

func init() {
	db = database.GetInstance()
}
