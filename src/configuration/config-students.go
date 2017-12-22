package configuration

import (
	"path/filepath"
)

type StudentConfig struct {
	IsReference bool
	Name        string `json:"name"`
	Path        string `json:"path"`
}

func loadStudentsConfig() {
	loadJsonConfig(studentsPath, &studentConfigs)
	for _, studentConfig := range studentConfigs {
		studentConfig.IsReference = false
		studentConfig.Path = filepath.Join(
			filepath.Dir(studentsPath), studentConfig.Path)
	}
	studentConfigs = append([]*StudentConfig{&StudentConfig{
		IsReference: true,
		Name:        "__REFERENCE_IMPLEMENTATION__",
	}}, studentConfigs...)
}
