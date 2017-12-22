package configuration

import (
	"github.com/ms-xy/Tutortool/src/database/models"
)

var (
	Tutor                   string
	ReferenceImpl           *models.Student
	Students                = map[uint]*models.Student{}
	StudentsList            []uint
	StudentsSortedModelList []*models.Student
	Tasks                   map[uint]*models.Task

	studentsPath   string
	studentConfigs = []*StudentConfig{}

	taskListPath string
	taskConfigs  = []*TaskConfig{}
)
