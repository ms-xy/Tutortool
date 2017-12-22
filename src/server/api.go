package server

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/ms-xy/Tutortool/src/server/api"
	"net/http"
	"strings"
	// logging
	"github.com/ms-xy/logtools"
	// traceback handling
	"fmt"
	"runtime"
	// "path/filepath"
	// "runtime/debug"
)

/*
Define functions for serving the API endpoints.
*/

func ServeAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		HandlePanic(w, recover())
	}()

	endpoint := strings.TrimLeft(ps.ByName("endpoint"), "/")

	logtools.WithFields(logtools.Fields{
		"endpoint": endpoint,
		"request":  r,
	}).Debug("Serving API")

	var err error

	if fn, ok := api.Handlers[endpoint]; ok {
		err = fn(w, r)
	} else {
		err = errors.New("Unknown API endpoint: " + endpoint)
	}

	// global handle for api errors - removes some extra code
	if err != nil {
		logtools.Errorf("%+v", err)
		http.Error(w, err.Error(), 500)
	}
}

func HandlePanic(w http.ResponseWriter, v interface{}) {
	if v != nil {
		if err, ok := v.(error); ok {
			if err != nil {
				logtools.WithFields(map[string]interface{}{
					"location": identifyPanic()}).Errorf("%+v", err)
				http.Error(w, err.Error(), 500)
				return
			}
		}
		logtools.WithFields(map[string]interface{}{
			"location": identifyPanic()}).Error(v)
		http.Error(w, fmt.Sprintf("%+v", v), 500)

	}
}

func identifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(5, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	// filepath.Base(name)
	return fmt.Sprintf("%s:%d", file, line)
}
