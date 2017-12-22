package testcase

import (
	// persistence layer
	"github.com/jinzhu/gorm"
	"github.com/ms-xy/Tutortool/src/database"

	// utility
	"errors"
	"fmt"
	"github.com/ms-xy/Tutortool/src/utility/request-tools"
	"net/http"

	// configuration
	"github.com/ms-xy/Tutortool/src/configuration"

	// http connection write helper
	"github.com/ms-xy/Tutortool/src/server/serverutils"
)

var (
	db *gorm.DB = database.GetInstance()
)

func GetTestcasesList(w http.ResponseWriter, r *http.Request) error {
	tid := requesttools.FormValueAs_uint(r, "taskID")
	task, exists := configuration.Tasks[tid]
	if !exists {
		return errors.New(fmt.Sprintf("task with id='%d' does not exist", tid))
	}
	return serverutils.WriteJSON(w, r, task.Testcases)
}

func Get(w http.ResponseWriter, r *http.Request) error {
	tid := requesttools.FormValueAs_uint(r, "taskID")
	num := requesttools.FormValueAs_int(r, "testcase")
	task, exists := configuration.Tasks[tid]
	if !exists {
		return errors.New(fmt.Sprintf("task with id='%d' does not exist", tid))
	}
	if num < 1 || num > len(task.Testcases) {
		return errors.New(fmt.Sprintf("testcase index out of bounds: '%d'", num))
	}
	return serverutils.WriteJSON(w, r, task.Testcases[num-1])
}

// func EvaluateReferenceImplementation(w http.ResponseWriter, r *http.Request) error {
// 	var (
// 		s_tcid    string
// 		tcid      int64
// 		testcase  *Testcase
// 		task      *Task
// 		gccresult *Result
// 		runresult *Result
// 		err       error
// 	)
// 	if s_tcid = r.FormValue("tcid"); s_tcid != "" {
// 		if tcid, err = strconv.ParseInt(s_tcid, 10, 64); err == nil {
// 			if testcase, err = db.GetTestcase(tcid); err == nil {
// 				if task, err = db.GetTask(testcase.Task); err == nil {
// 					// reference path is a fake submission path
// 					referencepath := filepath.Join(".", "tasks", task.Path)
// 					// student is a fake student
// 					student := &Student{ID: REFERENCE_IMPL, Name: "reference_impl", Results: []*Result{}}
// 					// run with fake params to get evaluation for the reference implementation
// 					gccresult, runresult, err = helpers.EvaluateSubmissionSingleTC(referencepath, task, testcase, student)
// 					// if execution was successful, remove the potentially existant old entry
// 					if err == nil {
// 						if err = db.RemoveResult(task.ID, testcase.ID, student.ID); err == nil {
// 							if err = db.AddResult(gccresult); err == nil {
// 								if runresult != nil {
// 									err = db.AddResult(runresult)
// 								}
// 							}
// 						}
// 					}
// 					// write out the result
// 					serverutils.WriteJSON(w, r, struct {
// 						Gccresult *Result `json:"gccresult"`
// 						Runresult *Result `json:"runresult"`
// 						Error     error   `json:"error"`
// 					}{gccresult, runresult, err})
// 				}
// 			}
// 		}
// 	} else {
// 		err = errors.New("No testcase ID supplied")
// 	}
// 	return err
// }
