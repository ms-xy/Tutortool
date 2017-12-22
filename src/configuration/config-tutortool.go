package configuration

type TutortoolConfig struct {
	Tutor          string `json:"tutor"`
	StudentsConfig string `json:"students-config"`
	TaskList       string `json:"task-list"`
}

func loadTutortoolConfig() {
	config := &TutortoolConfig{}
	loadJsonConfig("tutortool-config.json", config)
	Tutor = config.Tutor
	studentsPath = config.StudentsConfig
	taskListPath = config.TaskList
}
