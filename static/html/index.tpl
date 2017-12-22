{% extends "base.tpl" %}

{% block additional_css %}
    <style>
        .badge {
            float: none !important;
            margin-right: 10px;
        }
        #repetition-badge {
            margin: 0 5px;
        }
        .list-group {
            margin-bottom: 0px !important;
        }
    </style>
{% endblock %}

{% block content %}
<div class="page-header">
  <h1>Tutortool v2.0 Alpha</h1>
</div>

<div class="panel panel-primary">
    <div class="panel-heading"><b>Suggestion</b></div>
    <div class="panel-body">
        <ul class="list-group">
            <li class="list-group-item">
                <span class="badge">1</span>
                Use  the "Tasks" view to compile all submissions
            </li>
            <li class="list-group-item">
                <span class="badge">2</span>
                Use the "Tasks" view to run all submissions
            </li>
            <li class="list-group-item">
                <span class="badge">3</span>
                Use the "Students" view to examine the run result for a
                student's submission
            </li>
            <li class="list-group-item">
                <span class="badge">4</span>
                Use Moodle to assign a grade to the student
            </li>
            <li class="list-group-item">
                <span class="badge">5</span>
                Repeat from
                <span id="repetition-badge" class="badge">3</span>
                until all students are graded
            </li>
        </ul>
    </div>
</div>

<div class="panel panel-danger">
    <div class="panel-heading">
        <span class="glyphicon glyphicon-alert"></span>
        <b>Important Note</b>
    </div>
    <div class="panel-body">
        <p>
            Never close the browser tab whilst there are still operations
            pending.
            In case of doubt check with the server logs to see whether there
            are still actions executed.
        </p>
        <p>
            The client orchestrates task execution.
            It will send out each compilation or
            run task sequentially in order of queuing.
            Thus closing / reloading the tab will abort all those jobs.
        </p>
    </div>
</div>

<div class="panel panel-info">
    <div class="panel-heading"><b>"Students" View</b></div>
    <div class="panel-body">
        <p>
            In the Students view you may select a single student and compile or
            run their submission, as well as view the resulting output.
        </p>
        <p>
            The points shown in the students view have no meaning, currently.
            That's a feature planned for the future, but not even in the
            backlog currently.
        </p>
    </div>
</div>

<div class="panel panel-info">
    <div class="panel-heading"><b>Tasks</b></div>
    <div class="panel-body">
        The task view shows all configured tasks and their respective testcases.
        Additionally it exhibits buttons for each task that allow compiling and
        grading all students sequentially.
        <br>
        These functions may take a while before finishing.
    </div>
</div>

<div class="panel panel-info">
    <div class="panel-heading"><b>Tutortool Configuration</b></div>
    <div class="panel-body">
        <p>
            The Tutortool configuration is written in JSON and saved on the
            disk.
            It is not possible to change the configuration at runtime (yet).
        </p>
        <p>
            The main configuration file must lie next to the Tutortool
            executable and have the name tutortool-config.json. All other
            configuration file names and locations must be explicitely defined
            in that file.
        </p>
    </div>
</div>

{% endblock %}
