package configuration

import (
	"path/filepath"
	"time"
)

type TaskConfig struct {
	Name                    string            `json:"name"`
	Path                    string            `json:"path"`
	NamingSchema            string            `json:"naming-schema"`
	DueDate                 time.Time         `json:"due-date"`
	ReferenceImplementation string            `json:"reference-implementation"`
	Gcc                     *GccConfig        `json:"gcc"`
	Run                     *RunConfig        `json:"run"`
	Testcases               []*TestcaseConfig `json:"testcases"`
	Sources                 []string          `json:"sources"`
	// Sources is only important for the SourceView in the DetailView
}

type GccConfig struct {
	Parameters   []string          `json:"parameters"`
	Files        []string          `json:"files"`
	Replacements map[string]string `json:"replacements"`
}

type RunConfig struct {
	Timeout    int `json:"timeout"` // seconds
	StdoutSize int `json:"stdout-size"`
	StderrSize int `json:"stderr-size"`
}

type TestcaseConfig struct {
	Name       string   `json:"name"`
	Parameters []string `json:"parameters"`
	InputFile  string   `json:"input-file"`
	Points     int      `json:"points"`
	RunConfig
}

func loadTasksConfig() {
	// first load task list
	list := []string{}
	loadJsonConfig(taskListPath, &list)

	// based on this task list (index is the resulting taskID), create the
	// global Tasks  list - which in turn is persisted into the database
	taskConfigs = make([]*TaskConfig, len(list))
	for id, path := range list {
		taskConfigs[id] = loadTask(filepath.Join(filepath.Dir(taskListPath), path))
	}
}

func loadTask(path string) *TaskConfig {

	// set reasonable defaults
	// this is important, as it vastly reduces configuration file sizes

	taskConfig := &TaskConfig{
		NamingSchema:            "^hw(\\d+)$",
		DueDate:                 time.Now(),
		ReferenceImplementation: "reference_impl",
		Gcc: &GccConfig{
			Parameters: []string{
				"-Wall",
				"-Werror",
				"-Wno-deprecated-declarations",
				"-g",
				"-std=c99",
			},
			Files: []string{
				"*.c",
			},
		},
		Run: &RunConfig{
			Timeout:    180,
			StdoutSize: 10000,
			StderrSize: 10000,
		},
		Testcases: []*TestcaseConfig{},
		Sources:   []string{},
	}

	loadJsonConfig(path, taskConfig)
	taskConfig.Path = path
	return taskConfig
}
