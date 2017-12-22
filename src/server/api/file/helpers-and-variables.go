package file

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"errors"
	_errors "github.com/pkg/errors"
	"net/http"
	"strconv"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"
)

var (
	db *gorm.DB = database.GetInstance()

	ErrMissingParameter = errors.New("missing parameter")
	ErrNotFound         = errors.New("not found")
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func loadStudent(r *http.Request) (*models.Student, error) {
	if r == nil {
		return configuration.ReferenceImpl, nil
	}

	if s_sid := r.FormValue("studentID"); s_sid == "" {
		return nil, _errors.Wrap(ErrMissingParameter, "studentID")

	} else if sid, err := strconv.ParseUint(s_sid, 10, 32); err != nil {
		return nil, _errors.WithStack(err)

	} else if student, exists := configuration.Students[uint(sid)]; !exists {
		return nil, _errors.Wrapf(ErrNotFound, "student [id=%d]", uint(sid))

	} else {
		return student, nil
	}
}

func loadTask(r *http.Request) (*models.Task, error) {
	if s_tid := r.FormValue("taskID"); s_tid == "" {
		return nil, _errors.Wrap(ErrMissingParameter, "taskID")

	} else if tid, err := strconv.ParseUint(s_tid, 10, 32); err != nil {
		return nil, _errors.WithStack(err)

	} else if task, exists := configuration.Tasks[uint(tid)]; !exists {
		return nil, _errors.Wrapf(ErrNotFound, "task [id=%d]", uint(tid))

	} else {
		return task, nil
	}
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func stripPaths(paths []string, base string) []string {
	l := len(base)
	r := make([]string, len(paths))
	for i, path := range paths {
		if path[0:l] == base {
			r[i] = path[l:]
		} else {
			r[i] = path
		}
	}
	return r
}
