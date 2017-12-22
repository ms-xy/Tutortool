/*
Task list. Tasks are grouped in tables.
*/
class TaskListView {
  constructor ($parent) {
    this.$el = $("<table>", {
      addClass: "table",
      css: {
        "margin-bottom": "0px"
      }
    });
    $parent.append(this.$el);
  }

  add (taskView) {
    this.$el.append(taskView.$el);
  }

  destroy () {
    this.$el.remove();
  }
}

/*
View of a single task. Tasks are grouped in tables.
*/
class TaskView {
  constructor (task) {
    task.view = this;
    this.$prog = $("<td>", {
      css: {
        "width": "110px",
        "padding-right": "10px !important",
        "table-layout": "fixed",
        "background-repeat": "no-repeat"
      }
    });
    this.$name = $("<td>").html(task.name);
    this.$el = $("<tr>").append(this.$prog, this.$name);
  }

  setActiveTask () {
    this.$prog.css("background-image", "url('/static/images/progressing3.gif')")
  }

  destroy () {
    this.$el.remove();
  }
}

/*
Replacement task handler that ensures, that run and compile tasks are properly
displayed in the UI.
*/
class StudentUITaskHandler extends TaskHandler {
  constructor () {
    super();
    this.taskListView = new TaskListView($("#panel-tasks"));
  }

  enqueue(task) {
    console.log("TaskHandler.enqueue (",task,")");
    this._tasks.push(task);
    this.taskListView.add(new TaskView(task));
    this.run();
  }

  run() {
    if (!this._running) {
      this._running = true;
      let self = this;
      let runner = async () => {
        while(self._tasks.length > 0) {
          let task = self._tasks.shift();
          task.view.setActiveTask();
          try {
            await task._run();
          } catch (e) {
            if (e) {
              /*
              If e is undefined, then there was no error propagated, omit the
              error message in that case. It is most likely by means of Task
              anyways (Task._run())
              */
              console.error(e);
            }
          }
          task.view.destroy();
        }
        this._running = false;
      }
      runner();
    }
  }
}

// replace standard (silent) task handler by a task handler that graphically
// renders tasks to the UI
tutortool.api.setTaskHandler(new StudentUITaskHandler());
