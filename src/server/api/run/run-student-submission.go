package run

import (
	// persistence layer
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"github.com/ms-xy/Tutortool/src/utility/math-tools"
	"github.com/ms-xy/execute"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"

	// logging
	"github.com/ms-xy/logtools"
)

var (
	KiloByte = mathtools.PowUint64(10, 3)
	MegaByte = mathtools.PowUint64(10, 6)
	GigaByte = mathtools.PowUint64(10, 9)
)

func RunStudentSubmission(w http.ResponseWriter, r *http.Request) error {

	var (
		student    *models.Student
		task       *models.Task
		testcase   *models.Testcase
		taskResult *models.TaskResult
		result     *execute.ExecResult
		err        error
	)

	// load student
	if student, err = loadStudent(r); err != nil {
		return err
	}

	// load task
	if task, err = loadTask(r); err != nil {
		return err
	}

	// load testcase
	if testcase, err = loadTestcase(r, task); err != nil {
		return err
	}

	// get task result
	if taskResult, err = firstOrCreateTaskResult(student, task); err != nil {
		return err
	}

	// get and remove old runresult if exists
	newRunResults := make([]*models.RunResult, 0, len(taskResult.RunResults))
	for _, result := range taskResult.RunResults {
		if result.TestcaseNumber == testcase.Number {
			if err := db.Unscoped().Delete(result).Error; err != nil {
				return err
			}
		} else {
			newRunResults = append(newRunResults, result)
		}
	}
	taskResult.RunResults = newRunResults

	// get submission dir
	submissionDir := r.FormValue("submissionDir")
	if submissionDir == "" {
		return ErrEmptyParameterSubmissionDir
	}

	// assemble working directory
	workDir := filepath.Join(student.Path, submissionDir)

	// create rlimiter args list
	rlimiterArgs := []string{
		// set proc limit to 1 = disable forking/threading (including forkbombs)
		"-Hnproc", "1",
		// virtual and stack memory limit of 256MB each, 512MB total
		// this rather small limit should hopefully avoid host memory starvation
		"-Has", strconv.FormatUint(256*MegaByte, 10),
		//"-Hstack", strconv.FormatUint(256*MegaByte, 10),
		// disable core dumps = avoid file system cluttering
		"-Hcore", "0",
		// lower NICE boundary of 10 (20-limit) = avoid host cpu starvation
		"-Hnice", "10",
	}

	// run submission
	result, err = execute.Execute(&execute.Command{
		Executable: "a.out",
		WorkingDir: workDir,
		RlimitArgs: rlimiterArgs,
		Arguments:  testcase.Parameters,
		Input:      testcase.Input,
		Timeout:    taskOrTestcaseTimeout(task, testcase, 180*time.Second),
		StdoutSize: taskOrTestcaseStdoutSize(task, testcase, 10000),
		StderrSize: taskOrTestcaseStderrSize(task, testcase, 10000),
	})

	if err != nil {
		logtools.WithFields(map[string]interface{}{
			"error": err,
		}).Errorf("error executing %s's submission for task %d testcase %d",
			student.Name, task.ID, testcase.Number)
	}

	// create run result
	runresult := &models.RunResult{
		TaskResultID:   taskResult.ID,
		TestcaseNumber: testcase.Number,
		TestcaseID:     testcase.ID,
		ExecResult:     *result,
		Points:         0,
	}

	// store new run result
	taskResult.RunResults = append(taskResult.RunResults, runresult)
	if err = db.Create(runresult).Error; err != nil {
		logtools.Error(err)
		return err
	}

	// additionally return the result to the user
	return serverutils.WriteJSON(w, r, taskResult)
}
