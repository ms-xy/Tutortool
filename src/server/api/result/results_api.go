package result

import (
	// persistence layer
	"github.com/ms-xy/Tutortool/src/database/models"

	// utility
	"net/http"
	"strconv"

	// configuration

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"
)

func GetResultsList(w http.ResponseWriter, r *http.Request) error {
	var (
		student *models.Student
		err     error
	)

	if student, err = loadStudent(r); err != nil {
		return err
	}

	results := []*models.TaskResult{}
	err = db.
		Where("student_id=?", student.ID).
		Find(&results).
		Error

	if err != nil {
		return err
	}

	return serverutils.WriteJSON(w, r, results)
}

func Get(w http.ResponseWriter, r *http.Request) error {
	var (
		student *models.Student
		task    *models.Task
		err     error
	)

	if student, err = loadStudent(r); err != nil {
		return err
	}

	if task, err = loadTask(r); err != nil {
		return err
	}

	result := &models.TaskResult{}
	err = db.
		Where("student_id=? AND task_id=?", student.ID, task.ID).
		Find(result).
		Error

	if err != nil {
		return err
	}

	return serverutils.WriteJSON(w, r, result)
}

func GradeTaskResult(w http.ResponseWriter, r *http.Request) error {
	var (
		student *models.Student
		task    *models.Task
		err     error
	)

	if student, err = loadStudent(r); err != nil {
		return err
	}

	if task, err = loadTask(r); err != nil {
		return err
	}

	result := &models.TaskResult{}
	err = db.
		Where("student_id=? AND task_id=?", student.ID, task.ID).
		Find(result).
		Error

	if err != nil {
		return err
	}

	points, err := strconv.ParseInt(r.FormValue("points"), 10, 32)
	if err != nil {
		return err
	}

	result.Points = int(points)

	if err = db.Save(result).Error; err != nil {
		return err
	}

	return serverutils.WriteJSON(w, r, result)
}

func GradeTestcaseResult(w http.ResponseWriter, r *http.Request) error {
	var (
		student *models.Student
		task    *models.Task
		err     error
	)

	if student, err = loadStudent(r); err != nil {
		return err
	}

	if task, err = loadTask(r); err != nil {
		return err
	}

	result := &models.TaskResult{}

	err = db.
		Where("student_id=? AND task_id=?", student.ID, task.ID).
		Preload("RunResults").
		First(result).
		Error

	if err != nil {
		return err
	}

	testcaseID, err := strconv.ParseUint(r.FormValue("testcaseID"), 10, 32)
	if err != nil {
		return err
	}

	points, err := strconv.ParseInt(r.FormValue("points"), 10, 32)
	if err != nil {
		return err
	}

	for _, tcResult := range result.RunResults {
		if tcResult.TestcaseID == uint(testcaseID) {
			tcResult.Points = int(points)
			return serverutils.WriteJSON(w, r, tcResult)
		}
	}

	return ErrTestcaseNotFound
}
