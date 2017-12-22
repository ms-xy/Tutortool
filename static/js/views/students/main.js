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
  setTimeout(function(){createStudentTaskInfo()}, 0);
  initStudentNav();
}

/* -------------------------------------------------------------------------- */

function initStudentNav () {
  // get students as value list sorted by name
  studentList = tutortool.students.sort(function(a, b) {
    if (a.Name > b.Name) return 1;
    if (a.Name < b.Name) return -1;
    return 0;
  })

  var $ul = $("#nav-student"),
      $panel = $("#panel-student"),
      lastActiveID = tutortool.cache.get("views.students.nav.lastActive"),
      lastActive;

  for (let student of studentList) {
    // skip reference implementation
    if (student.IsReference) {
      continue;
    }

    // create nav entry
    let $a = $("<a>", {
      addClass: "student-selector",
      data: {"tutortool.student": student},
      text: student.Name,
      click: function (ev) {
        var $target = $(ev.target),
            student = $target.data("tutortool.student");

        tutortool.cache.store("views.students.nav.lastActive", student.ID)

        $(".student-selector").parent().removeClass("active");
        $target.parent().addClass("active");

        $(".student").hide();
        $(".student-"+student.ID).show();
      }
    });
    $ul.append(
      $("<li>").attr("role", "presentation").append($a)
    );
    if (student.ID == lastActiveID) {
      lastActive = $a;
    }

    // add info table, TODO: use Table class instead of manually constructing
    //                       the dom node.
    $panel.append(
      $("<div>", {
        addClass: "panel-body student student-"+student.ID,
        css: {"display": "none"},
      }).append(
        $("<table>").addClass("table").append($("<tbody>").append(
          $("<tr>").css({"background": "#ececff"}).append(
            $("<th>").attr("colspan", 2).text("Info")
          ),
          $("<tr>").append(
            $("<td>",{addClass: "col-xs-2", text: "Name:"}),
            $("<td>",{text: student.Name}),
          ),
          $("<tr>").append(
            $("<td>",{addClass: "col-xs-2", text: "Path:"}),
            $("<td>",{text: student.Path}),
          )
        )),
        $("<table>").addClass("table task-info-table").append(
          $("<tbody>").attr("id", "task-info-sid-"+student.ID)
        )
      )
    );
  }

  // enable nav
  if (lastActive) {
    lastActive.click();
  } else {
    $(".student-selector").first().click();
  }
}

function createStudentTaskInfo () {
  for (let entry of tutortool.students) {
    let student = entry[1];

    // skip reference implementation
    if (student.IsReference) {
      continue;
    }

    // otherwise create task info objects
    $table = $("#task-info-sid-"+student.ID);
    $table.append(
      $("<tr>").css({"background": "#ececff"}).append(
        $("<th>").text("Task-Name"),
        $("<th>"), // actions column
        $("<th>").text("Points"),
        $("<th>").text("Comment")
      )
    );
    for (let entry of tutortool.tasks) {
      let task = entry[1];
      let result = student.Results.get(task.ID);
      $table.append(
        $("<tr>").append(
          $("<td>").addClass("task-info-name").text(task.Name),
          $("<td>").addClass("task-info-actions").append(
            $("<div>").addClass("btn-group").attr("role", "group").append(
              buttonCompile(student, task, result),
              buttonRun(student, task, result),
              buttonDetails(student, task),
            ),
          ),
          $("<td>").addClass("task-info-points").text(taskInfoPoints(result)),
          $("<td>").addClass("task-info-comment").text(taskInfoComment(result))
        )
      );
    }
  }
}

/* -------------------------------------------------------------------------- */

function taskInfoPoints (result) {
  if (result) {
    return result.Points;
  }
  return 0;
}

function taskInfoComment (result) {
  if (result) {
    return result.Comment;
  }
  return "";
}

/* -------------------------------------------------------------------------- */

function buttonCompile (student, task, result) {
  var submissionDir = "hw"+(""+task.ID).padStart(2, '0');
  var button = new Button()
    .glyphicon("cog")
    .type("default")
    .tooltip("Compile")
    .click(function(ev){
      button.type("warning").$el.blur().attr("disabled", true).tooltip('hide');
      tutortool.api.compileStudentSubmission(student, task, submissionDir)
        .onStart(function(){
          button.type("primary").background("/static/images/progressing3.gif");
        })
        .onAlways(function(){
          button.background(undefined).$el.removeAttr("disabled");
        })
        .onFail(function(error){
          button.type("danger").tooltip("Server-Error: " + error);
        })
        .onDone(function(){
          button.type("success").tooltip("Compile")
        });
      });

  if (result && result.GccResult) {
    if (result.GccResult.Error != "") {
      button.type("danger").tooltip("Compile-Error: "+result.GccResult.Error)
    } else {
      button.type("success")
    }
  }

  return button.$el;
}

function buttonRun (student, task, result) {
  var submissionDir = "hw"+(""+task.ID).padStart(2, '0');
  var button = new Button()
    .glyphicon("play-circle")
    .type("default")
    .tooltip("Run")
    .click(function(ev){
      button.$el.blur().attr("disabled", true).tooltip('hide');
      tutortool.api.runStudentSubmission(student, task, submissionDir)
        .onStart(function(){
          button.type("primary").background("/static/images/progressing3.gif");
        })
        .onAlways(function(){
          button.background(undefined).$el.removeAttr("disabled");
        })
        .onFail(function(errors){
          button.type("danger").tooltip("Some testcases have errors");
        })
        .onDone(function(runResult){
          button.type("success").tooltip("Run");
        });
      });

  if (result && result.RunResults) {
    if (result.RunResults.length != task.Testcases.length) {
      button.type("warning").tooltip("Not all testcases have been run yet");

    } else {
      let hasErrors = false;
      for (let i=0; i<result.RunResults.length; i++) {
        if (result.RunResults[i].Error) {
          hasErrors = true;
        }
      }
      if (hasErrors) {
        button.type("danger").tooltip("Some testcases have errors");
      } else {
        button.type("success").tooltip("Run")
      }
    }
  }

  return button.$el;
}

function buttonDetails (student, task) {
  return new Button()
    .glyphicon("eye-open")
    .type("default")
    .tooltip("Show Details")
    .click(function(ev){
      ev.stopImmediatePropagation();
      $(ev.currentTarget).blur();
      (new DetailView(student, task)).show();
      return;

      // work the button
      let button = $(ev.currentTarget).data("tutortool.button"),
          lastButton = $(window).data("tutortool.buttons.active.detailview")

      if (lastButton) {
        lastButton.type("default");

        if (lastButton == button) {
          detailView.hide();
        }
      }
      button.type("primary").$el.blur();

      $(window).data("tutortool.buttons.active.detailview", button);

      // show data in detailview
      detailView.show(student, task);
    })
    .$el;
}
