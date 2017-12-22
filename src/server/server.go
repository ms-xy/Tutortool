package server

import (
	// server
	"github.com/julienschmidt/httprouter"
	"github.com/ms-xy/Tutortool/src/server/serverutils"
	"net/http"

	// go-bindata
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/ms-xy/Tutortool/src/bindata"

	// pongo2
	"github.com/flosch/pongo2"
	_ "github.com/ms-xy/Tutortool/src/server/filters"

	// utility
	"path/filepath"
	"strings"

	// logging
	"github.com/ms-xy/logtools"
)

var (
	fileserver = http.FileServer(&assetfs.AssetFS{Asset: bindata.Asset, AssetDir: bindata.AssetDir, AssetInfo: bindata.AssetInfo, Prefix: "/"})
	tpl_set    = pongo2.NewSet("gui", &GoBindataAssetsTemplateLoader{})
)

func init() {
	serverutils.Setup(tpl_set, getContext)
}

func ForwardIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, "/ui/index", 303)
}

func ServeTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	site := strings.Trim(strings.Replace(ps.ByName("site"), "/", "_", -1), "_")
	if site == "" {
		site = "index"
	}
	logtools.Debugf("Serving Template: %s", site)
	tpl, err := tpl_set.FromFile("static/html/" + site + ".tpl")
	if err != nil {
		logtools.Error(err.Error())
		http.NotFound(w, r)
	} else {
		if ctx, err := getContext(site, r); err == nil {
			err := tpl.ExecuteWriter(ctx, w)
			if err != nil {
				logtools.Error(err)
				serverutils.ServeError(w, r, err, 500)
			}
		} else {
			serverutils.ServeError(w, r, err, 500)
		}
	}
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// site := ps.ByName("site")
	// if site == "static" {
	path := ps.ByName("filepath")
	logtools.Debugf("Serving Static: %s", path)
	r.URL.Path = filepath.Clean("/static/" + path)
	fileserver.ServeHTTP(w, r)
	// } else {
	// http.NotFound(w, r)
	// }
}
