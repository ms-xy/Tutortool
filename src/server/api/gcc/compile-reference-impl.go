package gcc

import (
	// persistence layer
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"github.com/ms-xy/execute"
	"net/http"
	"path/filepath"

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"

	// logging
	"github.com/ms-xy/logtools"
)

func CompileReferenceImplementation(
	w http.ResponseWriter, r *http.Request) error {

	var (
		student    *models.Student
		task       *models.Task
		taskResult *models.TaskResult
		result     *execute.ExecResult
		err        error
	)

	// load student
	student, _ = loadStudent(nil)

	// load task
	if task, err = loadTask(r); err != nil {
		return err
	}

	// get task result
	if taskResult, err = firstOrCreateTaskResult(student, task); err != nil {
		return err
	}

	// remove gcc result if exists
	if taskResult.GccResult != nil {
		if err = db.Unscoped().Delete(taskResult.GccResult).Error; err != nil {
			return err
		}
	}

	// run gcc
	result, err = execute.Execute(&execute.Command{
		LookupPath: true,
		Executable: "gcc",
		WorkingDir: filepath.Join(filepath.Dir(task.Path), task.RefImplPath),
		Arguments:  createGccArguments(task),
		Input:      []byte{},
	})

	if err != nil {
		logtools.WithFields(map[string]interface{}{
			"error": err,
		}).Errorf("error compiling reference impl of task %d", task.ID)
	}

	// create gcc result
	gccresult := &models.GccResult{
		TaskResultID: taskResult.ID,
		ExecResult:   *result,
	}

	// store new gcc result
	taskResult.GccResult = gccresult
	if err = db.Create(gccresult).Error; err != nil {
		return err
	}

	// additionally return the result to the user
	return serverutils.WriteJSON(w, r, taskResult)
}
