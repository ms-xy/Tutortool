package server

import (
	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"

	// database
	// "github.com/ms-xy/Tutortool/src/database"
	// "github.com/ms-xy/Tutortool/src/database/models"

	// server
	"net/http"

	// html templating
	"github.com/flosch/pongo2"

	// utilities
	"errors"
	// "fmt"
	// "github.com/ms-xy/Tutortool/src/utility/request-tools"
	// logging
	// "github.com/ms-xy/logtools"
)

func getContext(site string, r *http.Request, attr ...interface{}) (pongo2.Context, error) {
	var err error = nil
	ctx := pongo2.Context{}
	switch site {

	case "students":
		err = fillContextStudents(ctx, r)

	case "error":
		if len(attr) == 2 {
			ctx["error_msg"] = attr[0]
			ctx["error_code"] = attr[1]
		} else {
			err = errors.New("Invalid arguments supplied for error view - need error_msg and error_code")
		}

	}
	ctx["view_"+site+"_active"] = "active"
	return ctx, err
}

func fillContextStudents(ctx pongo2.Context, r *http.Request) error {
	ctx["view"] = "overview"
	ctx["students"] = configuration.StudentsSortedModelList
	return nil
}
