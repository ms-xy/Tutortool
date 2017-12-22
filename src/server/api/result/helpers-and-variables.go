package result

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"errors"
	"net/http"
	"strconv"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"
)

var (
	db *gorm.DB = database.GetInstance()

	ErrEmptyParameterStudentID = errors.New("missing parameter studentID")
	ErrEmptyParameterTaskID    = errors.New("missing parameter taskID")

	ErrStudentNotFound  = errors.New("student not found")
	ErrTaskNotFound     = errors.New("task not found")
	ErrTestcaseNotFound = errors.New("testcase not found")
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func loadStudent(r *http.Request) (*models.Student, error) {
	if s_sid := r.FormValue("studentID"); s_sid == "" {
		return nil, ErrEmptyParameterStudentID

	} else if sid, err := strconv.ParseUint(s_sid, 10, 32); err != nil {
		return nil, err

	} else if student, exists := configuration.Students[uint(sid)]; !exists {
		return nil, ErrStudentNotFound

	} else {
		return student, nil
	}
}

func loadTask(r *http.Request) (*models.Task, error) {
	if s_tid := r.FormValue("taskID"); s_tid == "" {
		return nil, ErrEmptyParameterTaskID

	} else if tid, err := strconv.ParseUint(s_tid, 10, 32); err != nil {
		return nil, err

	} else if task, exists := configuration.Tasks[uint(tid)]; !exists {
		return nil, ErrTaskNotFound

	} else {
		return task, nil
	}
}
