package task

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"

	// utility
	"errors"
	"fmt"
	"github.com/ms-xy/Tutortool/src/utility/request-tools"
	"net/http"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"
)

var (
	db *gorm.DB = database.GetInstance()
)

func GetTasksList(w http.ResponseWriter, r *http.Request) error {
	return serverutils.WriteJSON(w, r, configuration.Tasks)
}

func Get(w http.ResponseWriter, r *http.Request) error {
	tid := requesttools.FormValueAs_uint(r, "id")
	task, exists := configuration.Tasks[tid]
	if !exists {
		return errors.New(fmt.Sprintf("task with id='%d' does not exist", tid))
	}
	return serverutils.WriteJSON(w, r, task)
}
