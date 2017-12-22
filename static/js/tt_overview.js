function ShowPointChart(anchor, sid) {
  // mark selected student in nav list as active
  $("#tt-studentslist .nav li").removeClass("active");
  $(anchor).parent().addClass("active");
  // fade out panels that need to be user activated
  $("#tt-taskview").addClass("hidden");
  $(".details-group").addClass("hidden");
  // grab data from the server
  tutortool.getTaskList(function(tasks){
    // console.log("tasks:",tasks);
    tutortool.cache.tasks = tasks;
    tutortool.getStudent(sid, function(student){
      // console.log("student:",student);
      tutortool.cache.student = student;
      tutortool.__showPointChart();
    })
  })
}

/* assign wrapper functions */
tutortool.ShowPointChart = ShowPointChart;

/* ---------------------------------------------------------------------------
helper functions
*/

function __showPointChart() {
  var $div = $("#tt-pointchart");
  $div.removeClass("hidden");
  // grab header and body
  var header = $("#pointchart-heading");
  var body   = $div.find(".panel-body");
  // set panel title
  header.text(tutortool.cache.student.name);

  // update button events
  // var navbar = body.find(".tt-panel-navbar");
  // navbar.find(".tt-button-refresh")
  //   .off("click")
  //   .on("click", function(){tutortool.ShowPointChart(details.Name)});

  // refine data for line chart
  var chart_labels = [""],
      chart_data = [{x:0,y:0}],
      i=0,
      result,
      points,
      maxPoints,
      results = {};
  for (;i<tutortool.cache.student.results.length;i++) {
    result = tutortool.cache.student.results[i];
    results[result.task] = result;
  }
  for (;i<tutortool.cache.tasks.length;i++) {
    if (!results[i]) {
      points = 0;
    } else {
      points = results[i].points;
    }
    maxPoints = tutortool.sumTaskPoints(tutortool.cache.tasks[i]);
    chart_data.push({
      x: i+1,
      y: points/maxPoints*100,
      task_id: tutortool.cache.tasks[i].id,
    });
    chart_labels.push(tutortool.cache.tasks[i].name);
  }

  // create line chart
  var $chart = body.find(".tt-panel-chart");
  if ($chart[0].chartjs) {
    var chart = $chart[0].chartjs
    chart.config.data.labels = chart_labels;
    chart.config.data.datasets[0].data = chart_data;
    chart.update();
  }
  else {
    var chart = new Chart($chart, {
      type: "line",
      data: {
        labels: chart_labels,
        datasets: [{
          label: "% of Points",
          data: chart_data,
          lineTension: 0,
          pointRadius: 10,
          pointHoverRadius: 15,
          borderColor: "#5577aa",
        }],
      },
      options: {
        scales: {
          xAxes: [{
            type: 'category',
            position: 'bottom'
          }],
          yAxes: [{
            type: 'linear',
            position: 'left',
            ticks: {
              fixedStepSize: 10,
              min: 0,
              max: 100,
              beginAtZero: true
            }
          }]
        },
        tooltips: {

        }
      }
    });
    $chart[0].chartjs = chart
  }

  // attach event handler to line chart
  $chart.off("click touchend").on("click touchend", function(e){
    var points = chart.getElementsAtEvent(e);
    if (points.length == 1) {
      var task_id = points[0]._chart.config.data.datasets[0].data[points[0]._index].task_id;
      $(".details-group").addClass("hidden");
      tutortool.__showTestcases(task_id);

      // $.get("/api/result/output", {student: tutortool.cache.student, task: taskID}).done(function(result){
      //   tutortool.cache.task_id     = taskID;
      //   tutortool.cache.task_result = data.TaskResult;
      // });
    }
  })
}

function __showTestcases(task_id) {
  // get the requested task
  var task = tutortool.cache.tasks.grepFirst(function(t){ return (t.id == task_id); })
  if (!task) {
    console.error("Oops, tried to load a non-existing task");
    return;
  }
  // create datasets for results and testcases, map key is the testcase ID
  var results = tutortool.cache.student.results
    .grep(function(r){ return (r.task == task_id); })
    .createMap(function(r){ return r.testcase; })
  var testcases = task.testcases
    .createMap(function(tc){ return tc.id; })
  // grab div and show it
  var $div = $("#tt-taskview").removeClass("hidden");
  // grab header and body
  var header = $("#taskview-heading");
  var body   = $div.find(".panel-body");
  // set panel title
  header.text(task.name);
  // clean up panel
  var container = body.find(".panel");
  container.find("tr").off("click touchend");
  container.children().remove();
  // create new table and add it to the panel
  var table = $("<table>")
    .addClass("table table-bordered table-striped table-hover")
    .append($("<thead>")
      .append($("<tr>")
        .append($("<th>").text("Testcase-Name"))
        .append($("<th>").text("Points"))
        // .append($("<th>"))
      )
    )
    .append($("<tbody>"));
  container.append(table);
  // populate table body with testcases (and the students results for these)
  var tbody = table.find("tbody").first();
  task.testcases.forEach(function(testcase, index) {
      var tcrp = "-";
      if (results[testcase.id]) {
        tcrp = results[testcase.id].points
      }
      tbody.append($("<tr>")
        .attr("tcid", testcase.id)
        .append($("<td>").text(testcase.name))
        .append($("<td>").text(tcrp+" / "+testcase.points))
        .on("click touchend", function(e){
          if (tutortool.cache.active_result) {
            tutortool.cache.active_result.removeClass("info");
          }
          var tr   = $(e.currentTarget).addClass("info"),
              tcid = tr.attr("tcid");
          tutortool.cache.active_result = tr;
          tutortool.__prepareShowResult(task, testcases[tcid], results[tcid]);
        })
      );
  });
}

function __prepareShowResult(task, testcase, result) {
  // populate reference impl cache for that task
  if (!task.refimpl_files) {
    tutortool.getTaskReferenceImpl(task.id, function(dirmap){
      if (!dirmap.exists) {
        console.warn("No reference implementation found for task id="+task.id);
        dirmap.files = [];
      }
      task.refimpl_exists = dirmap.exists;
      task.refimpl_files  = dirmap.files;
      tutortool.__prepareShowResult(task, testcase, result);
    });
    return;
  }
  // if the reference implementation has not been run yet / output is invalid
  // run it again
  if (!testcase.evaluated) {
    tutortool.evaluateReferenceImplementation(task.id, testcase.id, function(data){
      testcase.refimpl_result = data

      if (!data.gccresult) {
        console.error(data.gccresult);
        tutortool.showError("Compile Error", "Failed to compile the reference implementation<br/>gcc: unknown error, check local logfiles")

      } else if (data.gccresult.exec_exitcode != 0) {
        console.error(data.gccresult);
        tutortool.showError("Compile Error", "Failed to compile the reference implementation<br/>Exitcode: "+gccresult.exec_exitcode+"<code>"+gccresult.exec_stdout+"</code><code>"+gccresult.exec_stderr+"</code>)");

      } else if (!data.runresult) {
        tutortool.showError("Execute Error", "Failed to execute the reference implementation<br/>run: unknown error, check local logfiles")

      } else {
        testcase.evaluated = true;
        tutortool.__prepareShowResult(task, testcase, result);
      }
    });
    return;
  }
  // fetch students submission for this task (always anew, might change?)
  if (!tutortool.cache.student.submissions) {
    tutortool.cache.student.submissions = {};
  }
  if (!tutortool.cache.student.submissions[task.id]) {
    tutortool.getStudentSubmission(tutortool.cache.student.id, task.id, function(submission){
      if (submission.exists) {
        if (!submission.list) {
          tutortool.cache.student.submissions[task.id] = submission;
          tutortool.__prepareShowResult(task, testcase, result);
        } else {
          // TODO: display selection which submission to chose from
        }
      } else {
        // TODO: display notification that there is no suitable submission from this user
      }
    });
    return;
  }

  // if there is no result defined for this student-testcase pair, it needs to
  // be evaluated
  if (!result) {
    tutortool.evaluateStudentSubmission(task.id, testcase.id, tutortool.cache.student.id, function(data){
      console.log(data);
      // tutortool.__showResult(task, testcase, result);
    });
    return
  }
  tutortool.__showResult(task, testcase, result);
}

function __showResult(task, testcase, result) {
  // first of all un-hide the diffview
  $("#tt-diffview").removeClass("hidden");
  // Set title for views
  $("#diffview-heading").text("Result: "+testcase.name);
  // result can be undefined (not graded, which can additionally mean no
  // grading possible, i.e. if there's no submission for that task)
  var diffview = $("#diffview-body");
  if (result) {
    diffview.text("here goes the comparison between the outputs");
  } else {
    if (impl.exists) {
      if (impl.files && impl.files.length>0) {
        diffview.text("Student submission has not been graded yet");

      } else if (impl.list && impl.list.length>0) {
        diffview.text("There are multiple possible submissions to chose from:");
        // TODO: implement

      } else {
        tutortool.showError("Error Retrieving Student Submission", "Unexpected error happened, we haven't got a valid response. Please check local logs and report this issue.");
        return;
      }
    } else {
      diffview.text("Student did not hand in a submission for this task");
    }
  }
  // unhide the codeview, showing off all the files included in the students
  // submission as well as the files of the reference implementation
  var codeview = $("#tt-codeview").removeClass("hidden");
  var codeview_file = false;
  if (impl.exists) {
    if (impl.files) {
      impl.files.forEach(function(){
        console.log(arguments);
      })
    }
  }
  if (task.refimpl_exists) {
    if (task.refimpl_files) {
      task.refimpl_files.forEach(function(){
        console.log(arguments);
      })
    }
  }
}

/* assign helper functions */
tutortool.__showPointChart    = __showPointChart;
tutortool.__showTestcases     = __showTestcases;
tutortool.__prepareShowResult = __prepareShowResult;
tutortool.__ShowResult        = __showResult;
