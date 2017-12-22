package configuration

import (
	"github.com/ms-xy/Tutortool/src/database/datatypes"
	"github.com/ms-xy/Tutortool/src/database/models"
	"github.com/ms-xy/Tutortool/src/utility/assert"
	"github.com/ms-xy/Tutortool/src/utility/slices"

	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"fmt"
)

func createTestcaseRecords(task *models.Task, config *TaskConfig) {

	testcases := make([]*models.Testcase, len(config.Testcases))
	for number, config := range config.Testcases {
		// just as with tasks, start testcase ID at 1 - more natural
		testcases[number] = createTestcaseRecord(task, number+1, config)
	}

	for _, testcase := range testcases {
		assert.ErrorNil(db.Create(testcase).Error,
			fmt.Sprintf("error creating testcase record:\n%+v\n", testcase))
	}
	task.Testcases = testcases
}

func updateTestcaseRecords(task *models.Task, config *TaskConfig) {

	// prepare task.Testcases
	task.Testcases = make([]*models.Testcase, len(config.Testcases))

	// load existing testcase records
	testcases := []*models.Testcase{}
	assert.ErrorNil(
		db.Where("task_id=?", task.ID).Order("number ASC").Find(&testcases).Error,
		"error loading testcase records")

	// create id mappings
	num2id := make(map[int]uint, len(testcases))
	id2model := make(map[uint]*models.Testcase, len(testcases))
	id2used := make(map[uint]bool, len(testcases))

	for number, testcase := range testcases {
		num2id[number] = testcase.ID
		id2model[testcase.ID] = testcase
		id2used[testcase.ID] = false
	}

	// create/update step
	for number, config := range config.Testcases {
		var testcase *models.Testcase

		if id, exists := num2id[number]; !exists {
			// just as with tasks, start testcase ID at 1 - more natural
			testcase = createTestcaseRecord(task, number+1, config)
			assert.ErrorNil(db.Create(testcase).Error,
				"error creating testcase record")

		} else {
			testcase = id2model[id]
			updateTestcaseRecord(task, testcase, config)
			assert.ErrorNil(db.Save(testcase).Error,
				"error updating testcase record")
		}

		task.Testcases[number] = testcase
	}
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func createTestcaseRecord(
	task *models.Task, number int, config *TestcaseConfig,
) *models.Testcase {

	lastModified, inputBytes := loadInputFile(
		filepath.Dir(task.Path), config.InputFile)

	return &models.Testcase{
		TaskID: task.ID,

		Name:   config.Name,
		Number: number,

		Points:       config.Points,
		LastModified: lastModified,

		Input:      inputBytes,
		Parameters: datatypes.Strings(config.Parameters),
		Timeout:    time.Duration(config.Timeout) * time.Second,
		StdoutSize: config.StdoutSize,
		StderrSize: config.StderrSize,
	}
}

func updateTestcaseRecord(
	task *models.Task, testcase *models.Testcase, config *TestcaseConfig,
) {

	// same procedure as for tasks, if the input, points, or parameters changed
	// invalidate all related results

	testcase.Timeout = time.Duration(config.Timeout) * time.Second
	testcase.StdoutSize = config.StdoutSize
	testcase.StderrSize = config.StderrSize

	invalid := false

	if !slices.StringSlicesEqual(
		[]string(testcase.Parameters), config.Parameters) {

		invalid = true
		testcase.Parameters = datatypes.Strings(config.Parameters)
	}

	if testcase.Points != config.Points {
		invalid = true
		testcase.Points = config.Points
	}

	lastModified, inputBytes := loadInputFile(
		filepath.Dir(task.Path), config.InputFile)

	testcase.LastModified = lastModified

	if !bytes.Equal(testcase.Input, inputBytes) {
		invalid = true
		testcase.Input = inputBytes
	}

	// remove invalidated results
	if invalid {
		assert.ErrorNil(db.Exec("DELETE FROM run_results WHERE "+
			"task_result_id IN (SELECT id FROM task_results WHERE task_id=?) AND "+
			"testcase_id=?", task.ID, testcase.ID).Error,
			"error deleting invalidated run result records")
	}
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- //

func loadInputFile(taskPath, testcaseInputPath string) (time.Time, []byte) {
	if testcaseInputPath == "" {
		return time.Now(), []byte{}
	}

	path := filepath.Join(taskPath, testcaseInputPath)

	if fi, err := os.Stat(path); err != nil {
		panic("error accessing file: " + err.Error())
	} else {
		if fi.IsDir() {
			panic("error opening file: is directory")
		}
		lastModified := fi.ModTime()
		if fileData, err := ioutil.ReadFile(path); err != nil {
			panic("erro reading file: " + err.Error())
		} else {
			return lastModified, fileData
		}
	}
}
