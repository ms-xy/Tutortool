package gcc

import (
	// persistence layer
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"github.com/ms-xy/Tutortool/src/utility/fs"
	"github.com/ms-xy/execute"
	"net/http"
	"path/filepath"

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"

	// logging
	"github.com/ms-xy/logtools"
)

func CompileStudentSubmission(w http.ResponseWriter, r *http.Request) error {

	var (
		student    *models.Student
		task       *models.Task
		taskResult *models.TaskResult
		result     *execute.ExecResult
		err        error
	)

	// get student
	if student, err = loadStudent(r); err != nil {
		return err
	}

	// get task
	if task, err = loadTask(r); err != nil {
		return err
	}

	// get task result
	if taskResult, err =
		firstOrCreateTaskResult(student, task); err != nil {
		return err
	}

	// get submission dir
	submissionDir := r.FormValue("submissionDir")
	if submissionDir == "" {
		return ErrEmptyParameterSubmissionDir
	}

	// remove old gcc result if exists
	if taskResult.GccResult != nil {
		if err = db.Unscoped().Delete(taskResult.GccResult).Error; err != nil {
			return err
		}
	}

	// get working directory
	workDir := filepath.Join(student.Path, submissionDir)

	// copy all replacements from relative to the reference-impl into the student
	// directory, overwriting any files that reside there
	for replace, with := range task.GccReplacements {
		src := filepath.Join(filepath.Dir(task.Path), task.RefImplPath, with)
		dst := filepath.Join(workDir, replace)
		logtools.WithFields(logtools.Fields{
			"with":    src,
			"replace": dst,
		}).Warn("Replacing file in student submission")
		fs.CopyFile(src, dst)
	}

	// run gcc
	result, err = execute.Execute(&execute.Command{
		LookupPath: true,
		Executable: "gcc",
		WorkingDir: workDir,
		Arguments:  createGccArguments(task),
		Input:      []byte{},
	})

	if err != nil {
		logtools.WithFields(map[string]interface{}{
			"error": err,
		}).Errorf("error compiling student %s's submission of task %d",
			student.Name, task.ID)
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

	return serverutils.WriteJSON(w, r, taskResult)
}
