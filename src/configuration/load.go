package configuration

/*
Example directory structure at the location of execution is:
root/
  |---- tutortool (executable)
  |---- tutortool.sqlite3
  |---- tutortool-config.json
  |---- tasks/
  |       |---- tasks-config.json  // tutortool-config.json
  |       |---- task001/
  |       |---- task002/
  |       `---- ...
  |               |---- task-config.json  // tutortool-config.json
  |               |---- testcases/
  |               |       |---- testcase01 (files are read as binary and
  |               |       |                 directly passed as test input)
  |               |       |---- testcase02
  |               |       `---- ...
  |               `---- reference_impl/
  |                       |---- Makefile (used unless config specifies a compile
  |                       |               command)
  |                       `---- ...
  `---- students/
          |---- student1/
          |       |---- hw1/
          |       |       |---- Makefile
          |       |       `---- ....
          |       `---- ...
          |---- student2/
          |---- student3/
          `---- ...
*/

func Load() {
	loadTutortoolConfig()
	loadStudentsConfig()
	loadTasksConfig()

	updateDatabaseTasks()    // update tasks first, so we can preload results for
	updateDatabaseStudents() // students
}
