package run

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
	"time"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"
)

var (
	db *gorm.DB = database.GetInstance()

	ErrEmptyParameterStudentID     = errors.New("missing parameter studentID")
	ErrEmptyParameterTaskID        = errors.New("missing parameter taskID")
	ErrEmptyParameterSubmissionDir = errors.New("missing parameter submissionDir")
	ErrEmptyParameterTestcaseNum   = errors.New("missing parameter testcaseNum")

	ErrStudentNotFound  = errors.New("student not found")
	ErrTaskNotFound     = errors.New("task not found")
	ErrTestcaseNotFound = errors.New("testcase not found")
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

func loadTestcase(r *http.Request, task *models.Task) (*models.Testcase, error) {
	if s_tcnum := r.FormValue("testcaseNum"); s_tcnum == "" {
		return nil, ErrEmptyParameterTestcaseNum

	} else if tcnum, err := strconv.ParseInt(s_tcnum, 10, 32); err != nil {
		return nil, err

	} else if int(tcnum) > len(task.Testcases) {
		return nil, ErrTestcaseNotFound

	} else {
		// tcnum is offset by 1, as index starts with natural index 1
		return task.Testcases[int(tcnum)-1], nil
	}
}

func createGccArguments(task *models.Task) []string {
	return slices.StringSlicesJoin(
		task.GccParams, []string{"-o", "a.out"}, task.GccFiles)
}

// func firstOrCreateTaskResult(sid, tid uint) (*models.TaskResult, error) {
// 	taskResult := &models.TaskResult{
// 		StudentID:  sid,
// 		TaskID:     tid,
// 		Points:     0,
// 		Comment:    "",
// 		GccResult:  nil,
// 		RunResults: []*models.RunResult{},
// 	}
// 	err := db.
// 		Preload("RunResults").
// 		Where("student_id=? AND task_id=?", 0, tid).
// 		FirstOrCreate(taskResult).
// 		Error
// 	return taskResult, err
// }

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

func taskOrTestcaseTimeout(
	task *models.Task, testcase *models.Testcase, defaultTimeout time.Duration,
) time.Duration {

	if testcase.Timeout > 0 {
		return testcase.Timeout
	}
	if task.RunTimeout > 0 {
		return task.RunTimeout
	}
	return defaultTimeout
}

func taskOrTestcaseStdoutSize(
	task *models.Task, testcase *models.Testcase, defaultSize int) int {

	if testcase.StdoutSize > 0 {
		return testcase.StdoutSize
	}
	if task.RunStdoutSize > 0 {
		return task.RunStdoutSize
	}
	return defaultSize
}

func taskOrTestcaseStderrSize(
	task *models.Task, testcase *models.Testcase, defaultSize int) int {

	if testcase.StderrSize > 0 {
		return testcase.StderrSize
	}
	if task.RunStderrSize > 0 {
		return task.RunStderrSize
	}
	return defaultSize
}
