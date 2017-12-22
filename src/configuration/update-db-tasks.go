package configuration

import (
	"github.com/ms-xy/Tutortool/src/database/datatypes"
	"github.com/ms-xy/Tutortool/src/database/models"
	"github.com/ms-xy/Tutortool/src/utility/assert"
	"github.com/ms-xy/Tutortool/src/utility/slices"
	"github.com/ms-xy/logtools"

	"time"
)

func updateDatabaseTasks() {
	// get all currently registered tasks and map them by uid
	tasks := []*models.Task{}
	assert.ErrorNil(db.Find(&tasks).Error, "error loading task records")

	id2task := map[uint]*models.Task{}
	id2used := map[uint]bool{}
	for _, task := range tasks {
		id2task[task.ID] = task
		id2used[task.ID] = false
	}

	// iterate configured tasks, mark used, update existing, create missing
	for _id, config := range taskConfigs {
		id := uint(_id + 1) // start IDs at 1 instead of 0
		if task, exists := id2task[id]; !exists {
			task = createTaskRecord(config)
			id2task[id] = task
		} else {
			updateTaskRecord(task, config)
		}
		id2used[id] = true
	}

	// purge orphaned tasks
	for id, used := range id2used {
		if !used {
			deleteTaskRecord(id)
		}
	}

	// set global variable
	Tasks = id2task
}

func createTaskRecord(config *TaskConfig) *models.Task {
	task := &models.Task{
		Name:         config.Name,
		Path:         config.Path,
		NamingSchema: config.NamingSchema,
		DueDate:      config.DueDate,
		RefImplPath:  config.ReferenceImplementation,
		Testcases:    []*models.Testcase{},
	}

	task.GccParams = datatypes.Strings(config.Gcc.Parameters)
	task.GccFiles = datatypes.Strings(config.Gcc.Files)
	task.GccReplacements = datatypes.String2StringMap(config.Gcc.Replacements)

	task.RunTimeout = time.Duration(config.Run.Timeout) * time.Second
	task.RunStdoutSize = config.Run.StdoutSize
	task.RunStderrSize = config.Run.StderrSize

	task.Sources = datatypes.Strings(config.Sources)

	assert.ErrorNil(db.Create(task).Error,
		"error creating new task record")

	createTestcaseRecords(task, config)

	return task
}

func updateTaskRecord(task *models.Task, config *TaskConfig) {
	// slightly more complex than the student update
	// it's not enough to simply update the values:
	// if one of the gcc parameters, run parameters, or the input data changed
	// then it is necessary to re-evaluate all results for the task/testcase
	// (for all students obviously)
	// thus all affected results need to be marked Invalid=true

	invalid := false

	task.Path = config.Path
	task.Name = config.Name
	task.NamingSchema = config.NamingSchema
	task.DueDate = config.DueDate
	task.RefImplPath = config.ReferenceImplementation

	task.RunTimeout = time.Duration(config.Run.Timeout) * time.Second
	task.RunStdoutSize = config.Run.StdoutSize
	task.RunStderrSize = config.Run.StderrSize

	task.Sources = datatypes.Strings(config.Sources)

	// compare parameters
	if !slices.StringSlicesEqual(
		[]string(task.GccParams), config.Gcc.Parameters) {

		invalid = true
		logtools.Warn("parameters changed")
		logtools.Warnf("old=%+v ;; new=%+v", task.GccParams, config.Gcc.Parameters)
		task.GccParams = datatypes.Strings(config.Gcc.Parameters)
	}

	// compare files
	if !slices.StringSlicesEqual(
		[]string(task.GccFiles), config.Gcc.Files) {

		invalid = true
		logtools.Warn("files changed")
		logtools.Warnf("old=%+v ;; new=%+v", task.GccFiles, config.Gcc.Files)
		task.GccFiles = datatypes.Strings(config.Gcc.Files)
	}

	// compare replacements
	// now this is a bit more tricky, gotta get the key list and sort that one,
	// then compare the key list, then compare the values
	keysTask := make([]string, len(task.GccReplacements))
	keysConf := make([]string, len(config.Gcc.Replacements))
	i := 0
	for key, _ := range task.GccReplacements {
		keysTask[i] = key
		i++
	}
	i = 0
	for key, _ := range config.Gcc.Replacements {
		keysConf[i] = key
		i++
	}
	keysTask = slices.StringsMergesort(keysTask)
	keysConf = slices.StringsMergesort(keysConf)

	// check keys
	if !slices.StringSlicesEqual(keysTask, keysConf) {
		invalid = true
		logtools.Warn("replacement keys changed")
		task.GccReplacements = datatypes.String2StringMap(config.Gcc.Replacements)
	} else {
		// check values
		for _, key := range keysTask {
			if task.GccReplacements[key] != config.Gcc.Replacements[key] {
				invalid = true
				logtools.Warn("replacement values changed")
				task.GccReplacements =
					datatypes.String2StringMap(config.Gcc.Replacements)
				break
			}
		}
	}

	// save task
	assert.ErrorNil(db.Save(task).Error, "error updating task")

	// if anything important changed,
	if invalid {
		logtools.Warnf("Invalidating task results for task %+v", task)

		// 1) delete results associated with the task
		// - delete gcc result
		// - delete run results
		// - delete task results
		assert.ErrorNil(db.Exec("DELETE FROM gcc_results WHERE id IN "+
			"(SELECT id FROM task_results WHERE task_id=?)", task.ID).
			Error,
			"error removing invalidated gcc results")

		assert.ErrorNil(db.Exec("DELETE FROM run_results WHERE id IN "+
			"(SELECT id FROM task_results WHERE task_id=?)", task.ID).
			Error,
			"error removing invalidated run results")

		assert.ErrorNil(db.Exec("DELETE FROM task_results WHERE task_id=?",
			task.ID).Error,
			"error remvoing invalidated task results")
	}

	// create/update/delete testcase records for this task
	updateTestcaseRecords(task, config)
}

func deleteTaskRecord(id uint) {
	assert.ErrorNil(db.Exec("DELETE FROM gcc_results WHERE task_result_id"+
		" IN (SELECT id FROM task_results WHERE task_id=?)", id).Error,
		"error deleting orphaned gcc records")

	assert.ErrorNil(db.Exec("DELETE FROM run_results WHERE task_result_id"+
		" IN (SELECT id FROM task_results WHERE task_id=?)", id).Error,
		"error deleting orphaned run records")

	assert.ErrorNil(db.Exec("DELETE FROM task_results WHERE task_id=?", id).
		Error, "error deleting orphaned task result records")

	assert.ErrorNil(db.Exec("DELETE FROM testcases WHERE task_id=?", id).Error,
		"error deleting orphaned testcase records")

	assert.ErrorNil(db.Exec("DELETE FROM tasks WHERE id=?", id).Error,
		"error deleting task record")
}
