package api

import (
	// api endpoints
	"github.com/ms-xy/Tutortool/src/server/api/file"
	"github.com/ms-xy/Tutortool/src/server/api/gcc"
	"github.com/ms-xy/Tutortool/src/server/api/result"
	"github.com/ms-xy/Tutortool/src/server/api/run"
	"github.com/ms-xy/Tutortool/src/server/api/student"
	"github.com/ms-xy/Tutortool/src/server/api/task"
	"github.com/ms-xy/Tutortool/src/server/api/testcase"

	// utility
	"net/http"
)

var (
	Handlers = map[string]func(w http.ResponseWriter, r *http.Request) error{
		/*
			student handlers
		*/
		"students/list": student.GetStudentsList, //
		"student/get":   student.Get,             // studentID

		/*
			task handlers
		*/
		"tasks/list": task.GetTasksList, //
		"task/get":   task.Get,          // taskID

		/*
			testcase handlers
		*/
		"testcases/list": testcase.GetTestcasesList, // taskID
		"testcase/get":   testcase.Get,              // taskID, testcase (num)

		/*
			execution handlers
		*/
		"gcc/reference-implementation": gcc.CompileReferenceImplementation, // taskID
		"gcc/student-submission":       gcc.CompileStudentSubmission,       // studentID, taskID, submissionDir

		"run/reference-implementation": run.RunReferenceImplementation, // taskID, testcaseNum
		"run/student-submission":       run.RunStudentSubmission,       // studentID, taskID, testcaseNum, submissionDir

		/*
			Grading
		*/
		"results/list": result.GetResultsList, // studentID
		"result/get":   result.Get,            // studentID, taskID

		"result/grade/task":     result.GradeTaskResult,     // studentID, taskID, points
		"result/grade/testcase": result.GradeTestcaseResult, // studentID, taskID, testcaseID, points

		/*
			Files
		*/
		"file/glob":     file.GlobStudent,      // studentID, submissionDir, globpattern
		"file/ref/glob": file.GlobReference,    // taskID, globpattern
		"file/get":      file.GetFile,          // studentID, submissionDir, filename
		"file/ref/get":  file.GetReferenceFile, // taskID, filename
	}
)
