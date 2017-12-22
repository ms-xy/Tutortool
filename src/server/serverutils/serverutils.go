package serverutils

import (
	"encoding/json"
	pongo "github.com/flosch/pongo2"
	"net/http"
	// logging
	"github.com/ms-xy/logtools"
)

var (
	tpl_set    *pongo.TemplateSet
	getContext func(string, *http.Request, ...interface{}) (pongo.Context, error)
)

func Setup(_tpl_set *pongo.TemplateSet, _getContext func(string, *http.Request, ...interface{}) (pongo.Context, error)) {
	tpl_set = _tpl_set
	getContext = _getContext
}

func ServeError(w http.ResponseWriter, r *http.Request, err error, code int) {
	logtools.Error("Serving Error Page:", code, "-", err.Error())
	tpl, tplErr := tpl_set.FromFile("static/tpl/Error.tpl")
	if tplErr != nil {
		logtools.Error(tplErr.Error())
		http.NotFound(w, r)
	} else {
		if ctx, err := getContext("error", r, err.Error(), code); err == nil {
			tpl.ExecuteWriter(ctx, w)
		} else {
			logtools.Panic("Fatal Error:", err)
		}
	}
}

/*
Helper function to write JSON headers plus json body of some generic data
fragment. Returns an error on failure.
*/
func WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	var (
		bytes []byte
		err   error
	)
	if bytes, err = json.Marshal(data); err == nil {
		w.Header().Set("Content-Type", "text/json")
		_, err = w.Write(bytes)
	}
	return err
}
