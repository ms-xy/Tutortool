package database

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	// logging
	"github.com/ms-xy/logtools"
	// instance lock
	"sync"
	// datatypes
	// "github.com/ms-xy/Tutortool/src/database/models"
	// clean-up
	"os"
)

var (
	lock          = sync.Mutex{}
	instance      *gorm.DB
	test_instance *gorm.DB
)

func openDatabase(path string) *gorm.DB {
	// get instance
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		logtools.Error(err.Error())
		panic(err)
	}
	return db
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func GetInstance() *gorm.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = openDatabase("tutortool.sqlite3")
		}
	}
	return instance
}

func GetTestInstance() *gorm.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			// clear old database and open new one
			path := "test-database.sqlite3"
			os.Remove(path)
			instance = openDatabase(path)
		}
	}
	return instance
}
