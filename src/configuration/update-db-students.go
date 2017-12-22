package configuration

import (
	"github.com/ms-xy/Tutortool/src/database/models"
	"github.com/ms-xy/Tutortool/src/utility/assert"

	// logging
	"github.com/ms-xy/logtools"
)

func updateDatabaseStudents() {
	// at first load a list of all currently registered students
	students := []*models.Student{}
	assert.ErrorNil(db.
		Preload("Results").
		Preload("Results.GccResult").
		Preload("Results.RunResults").
		Find(&students).Error,
		"error loading student records")

	for _, student := range students {
		logtools.WithFields(logtools.Fields{
			"ID":   student.ID,
			"Name": student.Name,
			"Path": student.Path,
		}).Print("Student:")
	}

	// generate an name(string)->used(bool) mapping
	name2student := map[string]*models.Student{}
	path2student := map[string]*models.Student{}
	id2student := map[uint]*models.Student{}
	id2used := map[uint]bool{}
	for _, student := range students {
		name2student[student.Name] = student
		path2student[student.Path] = student
		id2student[student.ID] = student
		id2used[student.ID] = false
	}

	// go over all configured students, create new student models where necessary
	studentsList := make([]uint, 0, len(studentConfigs))
	studentsSortedModelsList := make([]*models.Student, 0, len(studentConfigs))

	for _, config := range studentConfigs {
		// either path or name must match to make it the same student
		student, exists := name2student[config.Name]
		if !exists {
			student, exists = path2student[config.Path]
		}
		// if it still does not exist, it must indeed be a new student
		if !exists {
			student = createStudentRecord(config)
			name2student[config.Name] = student
			id2student[student.ID] = student
		} else {
			updateStudentRecord(student, config)
		}
		id2used[student.ID] = true
		studentsList = append(studentsList, student.ID)
		studentsSortedModelsList = append(studentsSortedModelsList, student)
		if student.IsReference {
			ReferenceImpl = student
		}
	}

	// purge orphaned students
	for id, active := range id2used {
		if !active {
			deleteStudentRecord(id)
		}
	}

	// set global variables
	Students = id2student
	StudentsList = studentsList
	StudentsSortedModelList = studentsSortedModelsList
}

func createStudentRecord(config *StudentConfig) *models.Student {
	student := &models.Student{
		Name:        config.Name,
		Path:        config.Path,
		IsReference: config.IsReference,
	}
	assert.ErrorNil(db.Debug().Create(student).Error,
		"error creating new student record")
	return student
}

func updateStudentRecord(student *models.Student, config *StudentConfig) {
	student.Path = config.Path
	assert.ErrorNil(db.Save(student).Error,
		"error updating student record")
}

func deleteStudentRecord(id uint) {
	// 1) delete student task results
	// - delete gcc results
	// - delete run results
	// - delete task results
	assert.ErrorNil(db.Exec("DELETE FROM gcc_results WHERE task_result_id IN "+
		"(SELECT id FROM task_results WHERE student_id=?)", id).
		Error,
		"error removing orphaned gcc results")

	assert.ErrorNil(db.Exec("DELETE FROM run_results WHERE task_result_id IN "+
		"(SELECT id FROM task_results WHERE student_id=?)", id).
		Error,
		"error removing orphaned run results")

	assert.ErrorNil(db.Exec("DELETE FROM task_results WHERE student_id=?", id).
		Error,
		"error removing orphaned task results")

	// 2) delete student
	assert.ErrorNil(db.Exec("DELETE FROM students WHERE id=?", id).Error,
		"error deleting student record")
}
