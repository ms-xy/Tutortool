$(document).ready(function(){
  init();
});

async function init () {
  // load task definitions
  let taskList = await tutortool.api.getTasksList();
  tutortool.tasks = taskList.asMap(t => t.ID);

  // load student data, including results - this can take a while
  let studentList = await tutortool.api.getStudentsList();
  tutortool.students = studentList.asMap(s => s.ID);

  // mapify the task results for ease of access
  // also find reference impl and set on students object
  for (let entry of tutortool.students) {
    let student = entry[1];
    if (student.IsReference) {
      tutortool.students.referenceImpl = student;
    }
    if (!student.Results) {
      student.Results = new Map();
    } else {
      student.Results = student.Results.asMap(r => r.TaskID);
    }
  }

  // build page
  buildTaskList();

  // $("#crefimpl").click(function(){
  //   tutortool.api.compileReferenceImpl(tutortool.tasks.get(2)).onDone(r => {
  //     let stderr = r.GccResult.Stderr;
  //     console.log(stderr.decodeBase64Unicode());
  //   });
  // });
}

/* -------------------------------------------------------------------------- */

function buildTaskList () {
  let table = new Table().appendTo($("#panel-tasks")).append([
    new TableRow([
      new TableHeaderCell({text: "ID", addClass: "col-xs-1"}),
      new TableHeaderCell({text: "Name", addClass: "col-xs-3"}),
      new TableHeaderCell({text: "Status", addClass: "col-xs-6"}),
      new TableHeaderCell({text: "Actions", addClass: "col-xs-2"}),
    ]).css({
      "background-color": "#ececff"
    })
  ]);

  for (let entry of tutortool.tasks) {
    (new TaskView(entry[1])).appendTo(table);
  }
}

/*
TaskView
Consists of 2 rows, one is initially hidden and only revealed upon clicking onto
the visible row.
The TaskView consists of (a) the info-row, containing ID, name, status and
relevant action buttons, and (b) the detail-row, containing testcase information
and more detailed student status.
In order to provide the detail-row, it relies on content provided by the detail
view.
*/
class TaskView {
  constructor (task) {
    this.task = task;
    this.expanded = false;

    // TODO TODO: student view students page:
    // - fill in details in comparison to ref impl outputs

    // create the task row with status and actions
    this.row = new TableRow([
      new TableDataCell({text: ""+task.ID, addClass: "col-xs-1"}),
      new TableDataCell({text: task.Name, addClass: "col-xs-3"}),
      new TableDataCell({text: "-", addClass: "col-xs-6"}),
      new TableDataCell({cls: "col-xs-2"}).append((new ActionButtons(task)).$el)
    ]).click((ev) => {
      if (this.expanded) {
        this.row.removeClass("selected-row");
        this.details.hide();
        this.expanded = false;
      } else {
        this.row.addClass("selected-row");
        this.details.show();
        this.expanded = true;
      }
    });

    // create the actual detail view row, attach the detail view and hide it
    this.details = new TableRow([
      new TableDataCell().addClass("col-xs-1"),
      new TableDataCell().colspan(3).append((new TaskDetailView(task)).$el)
    ]).addClass("detail-row").hide();
  }
  appendTo (table) {
    this.row.appendTo(table);
    this.details.appendTo(table);
  }
}

/*
ActionButtons
This view provides actions for a task row, (a) compile all (including ref impl),
(b) run all (including ref impl).
It heavily relies on the Button class.
*/
class ActionButtons {
  constructor (task) {
    this.$el = $("<div>").addClass("btn-group").attr({"role": "group"});

    // submission dir won't ever change for that task
    let submissionDir = "hw"+(""+task.ID).padStart(2, '0');

    // build the compile button
    let compile = new Button()
      .glyphicon("cog")
      .type("default")
      .tooltip("Compile All")
      .addClass("btn-sm")
      .appendTo(this.$el)
      .click(function(ev){

        // don't trigger row toggle
        ev.stopImmediatePropagation();

        // update visuals
        compile
          .type("warning")
          .$el
            .blur()
            .attr("disabled", true)
            .tooltip('hide');

        // execute queries for every student using api tasks
        let i=0, hasErrors=false;
        for (let entry of tutortool.students) {

          // local variables
          let student = entry[1],
              isFirst = (i == 0),
              isLast  = (i == (tutortool.students.size-1)),
              apiCall;

          // increment the counter
          i++;

          // special treatment for the ref impl
          if (student.IsReference) {
            apiCall = tutortool.api.compileReferenceImpl(task);
          } else {
            apiCall = tutortool.api.compileStudentSubmission(
              student, task, submissionDir
            );
          }

          apiCall
            .onStart(() => {
              // update visuals to add the progress bar if it is the first query
              if (isFirst) {
                compile
                  .type("primary")
                  .background("/static/images/progressing3.gif");
              }
            })
            .onAlways(() => {
              // update visuals to remove the progress bar and enable the button
              // if it is the last query
              if (isLast) {
                compile
                  .background(undefined)
                  .$el
                    .removeAttr("disabled");
              }
            })
            .onFail((error) => {
              // set hasErrors flag to avoid setting the button to green
              // accidentally in the done callback of another successful
              // compiler run
              hasErrors = true;
              if (isLast) {
                compile
                  .type("danger")
                  .tooltip(error);
              }
            })
            .onDone(() => {
              // update visuals to indicate success if all compilations ran
              if (isLast) {
                if (!hasErrors) {
                  compile
                    .type("success")
                    .tooltip("Compile All");
                } else {
                  compile
                    .type("danger")
                    .tooltip("Some submissions could not be compiled");
                }
              }
            });
        }
      });

    // build the run button
    let run = new Button()
      .glyphicon("play-circle")
      .type("default")
      .tooltip("Run")
      .addClass("btn-sm")
      .appendTo(this.$el)
      .click(function(ev){

        // don't trigger row toggle
        ev.stopImmediatePropagation();

        // update visuals
        run
          .type("warning")
          .$el
            .blur()
            .attr("disabled", true)
            .tooltip('hide');

        // execute queries for every student using api tasks
        let i=0, hasErrors=false;
        for (let entry of tutortool.students) {

          // local variables
          let student = entry[1],
              isFirst = (i == 0),
              isLast  = (i == (tutortool.students.size-1)),
              apiCall;
          i++;

          if (student.IsReference) {
            apiCall = tutortool.api.runReferenceImpl(task);
          } else {
            apiCall = tutortool.api.runStudentSubmission(
              student, task, submissionDir
            );
          }

          apiCall
            .onStart(() => {
              // update visuals to show the progress if is the first query
              if (isFirst) {
                run
                  .type("primary")
                  .background("/static/images/progressing3.gif");
              }
            })
            .onAlways(() => {
              // update visuals, remove progress and enable button if last
              if (isLast) {
                run
                  .background(undefined)
                  .$el
                    .removeAttr("disabled");
              }
            })
            .onFail((errors) => {
              hasErrors = true;
              if (isLast) {
                run
                  .type("danger")
                  .tooltip("Some submissions could not be run successfully");
              }
            })
            .onDone(() => {
              // update visuals, only indicate success if there were no errors
              if (isLast) {
                if (!hasErrors) {
                  run
                    .type("success")
                    .tooltip("Run All");
                } else {
                  run
                    .type("danger")
                  .tooltip("Some submissions could not be run successfully");
                }
              }
            });
        }
      });
  }
}

/*
TaskDetailView
This element provides more in-depth details about a given task.
It provides a bootstrap 'well' element as the root element, in which it groups
two tables, (a) the testcase table, showing information about testcases,
and (b) the student table, showing information about students.
It relies on the TestcaseView and StudentView to provide its information.
*/
class TaskDetailView {
  constructor (task) {
    // prepare elements for the detail view
    this.$el = $("<div>").addClass("well well-sm");

    let testcasePanel = $("<div>")
          .addClass("panel panel-default")
          .appendTo(this.$el)
          .append($("<div>")
            .addClass("panel-heading")
            .text("Show Testcases")
            .click((ev) => {
              $(ev.currentTarget).hide().next().show();
            })),
        testcaseTable = new Table(testcasePanel).hide(),
        studentTable = new Table(testcasePanel).hide();

    // setup testcase table and click action
    testcaseTable.append(new TableRow([
      new TableHeaderCell({text: "Number", addClass: "col-xs-1"}),
      new TableHeaderCell({text: "Name", addClass: "col-xs-3"})
    ]).css({
      "background-color": "#ececff"
    }).click((ev) => {
      $(ev.currentTarget).closest("table").hide().prev().show();
    }))

    // create testcase views for each testcase and append to testcase table
    for (let testcase of task.Testcases) {
      (new TestcaseView(testcase)).appendTo(testcaseTable);
    }
  }
}

class TestcaseView {
  constructor (testcase) {
    this.testcase = testcase;
    this.row = new TableRow([
      new TableDataCell({text: ""+testcase.Number, addClass: "col-xs-1"}),
      new TableDataCell({text: testcase.Name, addClass: "col-xs-3"}),
    ]);
    // TODO: action buttons (compile all + run all)
  }
  appendTo (table) {
    this.row.appendTo(table);
  }
}

/* -------------------------------------------------------------------------- */
