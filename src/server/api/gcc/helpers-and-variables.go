package gcc

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"errors"
	"github.com/ms-xy/Tutortool/src/utility/slices"
	"net/http"
	"strconv"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"
)

var (
	db *gorm.DB = database.GetInstance()

	ErrEmptyParameterStudentID     = errors.New("missing parameter studentID")
	ErrEmptyParameterTaskID        = errors.New("missing parameter taskID")
	ErrEmptyParameterSubmissionDir = errors.New("missing parameter submissionDir")

	ErrStudentNotFound = errors.New("student not found")
	ErrTaskNotFound    = errors.New("task not found")
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func loadStudent(r *http.Request) (*models.Student, error) {
	if r == nil {
		return configuration.ReferenceImpl, nil
	}

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

func createGccArguments(task *models.Task) []string {
	return slices.StringSlicesJoin(
		task.GccParams, []string{"-o", "a.out"}, task.GccFiles)
}

func firstOrCreateTaskResult(
	student *models.Student, task *models.Task) (*models.TaskResult, error) {

	for _, result := range student.Results {
		if result.TaskID == task.ID {
			return result, nil
		}
	}

	taskResult := &models.TaskResult{
		StudentID:  student.ID,
		TaskID:     task.ID,
		Points:     0,
		Comment:    "",
		GccResult:  nil,
		RunResults: []*models.RunResult{},
	}

	student.Results = append(student.Results, taskResult)

	return taskResult, db.Create(taskResult).Error
}
