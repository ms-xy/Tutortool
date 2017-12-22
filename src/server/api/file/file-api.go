package file

import (
	// persistence layer
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"github.com/ms-xy/Tutortool/src/server/serverutils"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func GlobStudent(w http.ResponseWriter, r *http.Request) error {

	var (
		student *models.Student
		err     error
	)

	// get student
	if student, err = loadStudent(r); err != nil {
		return errors.WithStack(err)
	}

	// get submission dir
	submissionDir := r.FormValue("submissionDir")
	if submissionDir == "" {
		return errors.Wrap(ErrMissingParameter, "submissionDir")
	}

	// get globpattern
	globpattern := r.FormValue("globpattern")
	if globpattern == "" {
		return errors.Wrap(ErrMissingParameter, "globpattern")
	}

	// assemble file path
	basepath := filepath.Join(student.Path, submissionDir)
	globpath := filepath.Join(basepath, globpattern)

	// use globbing to determine filepaths
	resultpaths, err := filepath.Glob(globpath)
	if err != nil {
		return errors.WithStack(err)
	}

	// strip basepath of results
	resultpaths = stripPaths(resultpaths, basepath)

	// write result
	return serverutils.WriteJSON(w, r, resultpaths)
}

func GlobReference(w http.ResponseWriter, r *http.Request) error {

	var (
		task *models.Task
		err  error
	)

	// get task
	if task, err = loadTask(r); err != nil {
		return errors.WithStack(err)
	}

	// get globpattern
	globpattern := r.FormValue("globpattern")
	if globpattern == "" {
		return errors.Wrap(ErrMissingParameter, "globpattern")
	}

	// assemble file path
	basepath := filepath.Join(filepath.Dir(task.Path), task.RefImplPath)
	globpath := filepath.Join(basepath, globpattern)

	// use globbing to determine filepaths
	resultpaths, err := filepath.Glob(globpath)
	if err != nil {
		return errors.WithStack(err)
	}

	// strip basepath of results
	resultpaths = stripPaths(resultpaths, basepath)

	// write result
	return serverutils.WriteJSON(w, r, resultpaths)
}

func GetFile(w http.ResponseWriter, r *http.Request) error {

	var (
		student *models.Student
		err     error
	)

	// get student
	if student, err = loadStudent(r); err != nil {
		return errors.WithStack(err)
	}

	// get submission dir
	submissionDir := r.FormValue("submissionDir")
	if submissionDir == "" {
		return errors.Wrap(ErrMissingParameter, "submissionDir")
	}

	// get filename
	filename := r.FormValue("filename")
	if filename == "" {
		return errors.Wrap(ErrMissingParameter, "filename")
	}

	// assemble file path
	fpath := filepath.Join(student.Path, submissionDir, filename)

	// open file and write result
	if bytes, err := ioutil.ReadFile(fpath); err != nil {
		return errors.WithStack(err)
	} else {
		if _, err := w.Write(bytes); err != nil {
			return errors.WithStack(err)
		} else {
			w.WriteHeader(200)
			return nil
		}
	}
}

func GetReferenceFile(w http.ResponseWriter, r *http.Request) error {

	var (
		task *models.Task
		err  error
	)

	// get task
	if task, err = loadTask(r); err != nil {
		return errors.WithStack(err)
	}

	// get filename
	filename := r.FormValue("filename")
	if filename == "" {
		return errors.Wrap(ErrMissingParameter, "filename")
	}

	// assemble file path
	fpath := filepath.Join(filepath.Dir(task.Path), task.RefImplPath, filename)

	// open file and write result
	if bytes, err := ioutil.ReadFile(fpath); err != nil {
		return errors.WithStack(err)
	} else {
		if _, err := w.Write(bytes); err != nil {
			return errors.WithStack(err)
		} else {
			w.WriteHeader(200)
			return nil
		}
	}
}
