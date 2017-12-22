package main

import (
	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"

	// server
	"github.com/julienschmidt/httprouter"
	"github.com/ms-xy/Tutortool/src/server"
	"net/http"

	// logging
	"github.com/ms-xy/logtools"

	_ "net/http/pprof"
)

func main() {
	// Scan directories for tasks, testcases and unregistered students
	configuration.Load()

	// Router and Mux for requests
	mux := http.NewServeMux()
	router := &httprouter.Router{}

	// Serve http files
	router.GET("/", server.ForwardIndex)
	router.GET("/ui/*site", server.ServeTemplate)
	router.GET("/static/*filepath", server.ServeStaticFiles)

	// Serve api calls
	router.POST("/api/*endpoint", server.ServeAPI)
	router.GET("/api/*endpoint", server.ServeAPI)

	// launch pprof
	go func() {
		logtools.Info("Binding pprof to localhost:6060")
		logtools.Info(http.ListenAndServe("localhost:6060", nil))
	}()

	// Launch http file server on port 8080
	mux.Handle("/", router)
	logtools.Info("Binding to :8080")
	logtools.Fatal(http.ListenAndServe(":8080", mux))
}
