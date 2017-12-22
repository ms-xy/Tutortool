/*
Evidentally we need a more differenciated view
*/
class DetailView {

  constructor (student, task) {
    let gccRow    = new GccView(student, task),
        runRow    = new RunView(student, task),
        sourceRow = new SourceView(student, task),
        headerRow = (new Row()).append(
          (new Column(6)).append(
            (new Panel()).heading.text("Reference Implementation").$el.parent()
          ).$el,
          (new Column(6)).append(
            (new Panel()).heading.text(student.Name).$el.parent()
          ).$el
        );

    this.viewport = new Viewport(2000,
      $("<div>", {addClass: "container"}).append([
        $("<h3></h3>").text(task.Name),
        headerRow.$el,
        gccRow.$el,
        runRow.$el,
        sourceRow.$el
      ])
    ).on("hide", () => {this.viewport.destroy()});
  }

  show () {
    this.viewport.show();
    return this;
  }

  hide () {
    this.viewport.destroy();
    return this;
  }
}

class GccView {
  constructor (student, task) {

    // create dom structure
    let row      = new Row(),
        leftCol  = new Column(6, row.$el),
        rightCol = new Column(6, row.$el),

        leftPanel = new Panel("default", leftCol.$el),
        rightPanel = new Panel("default", rightCol.$el);

    this.$el = row.$el;

    // access to data
    let referenceImpl       = tutortool.students.referenceImpl,
        referenceTaskResult = referenceImpl.Results.get(task.ID),
        studentTaskResult   = student.Results.get(task.ID);

    // if either the result is missing or no gcc-result is available, display
    // a message indicating so
    // otherwise display a table with the details regarding compilation
    if (!referenceTaskResult || !referenceTaskResult.GccResult) {
      leftPanel.body.text(
        "Please compile the reference implementation first")
    } else {
      leftPanel.append(makeGccTable(referenceTaskResult.GccResult));
    }

    // same for student result
    if (!studentTaskResult || !studentTaskResult.GccResult) {
      rightPanel.body.text("Please compile the student submission first")
    } else {
      rightPanel.append(makeGccTable(studentTaskResult.GccResult));
    }
  }
}

class RunView {
  constructor (student, task) {
    this.$el = $();

    // get results
    let referenceImpl = tutortool.students.referenceImpl,
        rtr = referenceImpl.Results.get(task.ID),
        str = student.Results.get(task.ID),
        rtrOk = false,
        strOk = false,
        rtcResults, stcResults;

    if (rtr && rtr.RunResults) {
      rtrOk = true;
      rtcResults = rtr.RunResults.asMap((r) => r.TestcaseNumber);
    }

    if (str && str.RunResults) {
      strOk = true;
      stcResults = str.RunResults.asMap((r) => r.TestcaseNumber);
    }

    for (let testcase of task.Testcases) {

      // create dom structure
      let row = new Row(),
          leftCol  = new Column(6, row.$el),
          rightCol = new Column(6, row.$el),

          leftPanel = new Panel("default", leftCol.$el),
          rightPanel = new Panel("default", rightCol.$el);

      this.$el = this.$el.add(row.$el);

      // get testcase results if available
      let rtcResult=undefined, stcResult=undefined;
      if (rtrOk) {
        rtcResult = rtcResults.get(testcase.Number);
      }
      if (strOk) {
        stcResult = stcResults.get(testcase.Number);
      }

      // if the task result is missing or if the specific testcase result is
      // missing, render an info message
      if (!rtcResult) {
        leftPanel.body.text(
          "Please execute the reference implementation first")
      } else {
        leftPanel.append(makeRunTable(testcase, rtcResult));
      }

      // same for student result
      if (!stcResult) {
        rightPanel.body.text("Please execute the student submission first")
      } else {
        rightPanel.append(makeRunTable(testcase, stcResult));
      }
    }
  }
}

class SourceView {
  constructor (student, task) {
    this.$el = $("<div>"); //.addClass("row");
    setTimeout(()=>{
      this.load(student, task, task.Sources)
    }, 0);

    //   // create dom structure
    //   let row = new Row(),
    //       leftCol  = new Column(6, row.$el),
    //       rightCol = new Column(6, row.$el),

    //       leftPanel = new Panel("default", leftCol.$el),
    //       rightPanel = new Panel("default", rightCol.$el);

    //   this.$el = this.$el.add(row.$el);
  }

  async load (student, task, files) {
    var dir = "hw"+(""+task.ID).padStart(2, '0');
    for (let globpattern of files) {
      let referenceFilenames = await tutortool.api.globReferenceFiles(task.ID, globpattern);
      let studentFilenames = await tutortool.api.globStudentFiles(student.ID, dir, globpattern);
      let rFiles = [];
      for (let filename of referenceFilenames) {
        rFiles.push({
          name: filename,
          data: await tutortool.api.getReferenceFile(task.ID, filename)
        });
      }
      let sFiles = [];
      for (let filename of studentFilenames) {
        sFiles.push({
          name: filename,
          data: await tutortool.api.getFile(student.ID, dir, filename)
        });
      }
      this.makeRow(rFiles, sFiles);
    }
  }

  makeRow (referenceFiles, studentFiles) {
      // create dom structure
      let row = new Row(),
          leftCol  = new Column(6, row.$el),
          rightCol = new Column(6, row.$el);

      for (let file of referenceFiles) {
        let panel = new Panel("default", leftCol.$el)
              .addClass("x-small-text"),
            pre = new PreformattedText()
              .addClass("source-file-listing")
              .html(Prism.highlight(file.data, Prism.languages.c));
        panel.heading.text(file.name);
        panel.body.append(pre.$el);
      }

      for (let file of studentFiles) {
        let panel = new Panel("default", rightCol.$el)
              .addClass("x-small-text"),
            pre = new PreformattedText()
              .addClass("source-file-listing")
              .html(Prism.highlight(file.data, Prism.languages.c));
        panel.heading.text(file.name);
        panel.body.append(pre.$el);
      }

      this.$el.append(row.$el);
  }
}

class X {
  _create_RunResults (runResults, tasks) {
    if (!runResults) {
      return;
    }

    for (let runResult of runResults) {
      this.$el.append($("<table>").addClass("table").append(
        $("<tbody>").append(
          $("<tr>").css("background-color", "rgb(236, 236, 255)").append(
            $("<th colspan=2>").text("Testcase: "+runResult.TestcaseNumber)
          ),
          $("<tr>").append(
            $("<td>").text("ExitCode"),
            $("<td>").text(""+runResult.ExitCode)
          ),
          $("<tr>").append(
            $("<td>").text("Error"),
            $("<td>").text(runResult.Error)
          ),
          $("<tr>").append(
            $("<td>").text("KillReason (if killed)"),
            $("<td>").text(runResult.KillReason)
          ),
          $("<tr>").append(
            $("<td>").text("Stdout"),
            $("<td>").append(
              $("<pre>").text(
                (runResult.Stdout)?atob(runResult.Stdout).substr(0,1000):""))
          ),
          $("<tr>").append(
            $("<td>").text("Stderr"),
            $("<td>").append(
              $("<pre>").text(
                (runResult.Stderr)?atob(runResult.Stderr).substr(0,1000):""))
          )
        )
      ))
    }
  }

  show (student, task) {
    this.$el.children().remove();

    this._create_title(student, task);

    if (!student.Results) {
      student.Results = [];
    }

    let taskResult = undefined;
    for (let _taskResult of student.Results) {
      if (_taskResult.TaskID == task.ID) {
        taskResult = _taskResult;
      }
    }

    if (!taskResult) {
      this._create_EmptyResultMsg();
    } else {
      this._create_GccResult(taskResult.GccResult);
      this._create_RunResults(taskResult.RunResults, task);
    }
  }

  hide () {
    this.$el.children().remove();
    this.$el.hide();
  }
}