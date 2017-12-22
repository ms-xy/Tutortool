package student

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

func GetStudentsList(w http.ResponseWriter, r *http.Request) error {
	return serverutils.WriteJSON(w, r, configuration.Students)
}

func Get(w http.ResponseWriter, r *http.Request) error {
	sid := requesttools.FormValueAs_uint(r, "id")
	student, exists := configuration.Students[sid]
	if !exists {
		return errors.New(fmt.Sprintf("student with id='%d' does not exist", sid))
	}
	return serverutils.WriteJSON(w, r, student)
}
