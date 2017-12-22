{% extends "base.tpl" %}

{% block additional_css %}
    <link href="/static/css/tt_overview.css" rel="stylesheet" />
    <!--
    Include prismjs for syntax highlighting of source codes
    (students submissions)
    -->
    <link href="/static/css/prism.min.css" rel="stylesheet" />
{% endblock %}

{% block additional_js %}
    <script src="/static/js/tt_overview.js"></script>
    <!--
    Include prismjs for syntax highlighting of source codes
    (students submissions)
    Make sure it has the data-manual attribute so it does not attempt to
    automagically format code
    -->
    <script src="/static/js/prism.min.js" data-manual></script>
    <!--
    Include diff.js (by kpdecker (see github)) for advanced text diff
    -->
    <script src="/static/js/diff.min.js"></script>
{% endblock %}

{% block content %}
    <div class='row'>
        <div id='tt-studentslist' class='{{col3}}'>
            <div class='panel panel-default'>
                <div class='panel-body'>
                    <ul class='nav nav-pills nav-stacked'>
                        {% if not students %}
                            <p class='alert alert-danger'>No students found</p>
                        {% endif %}
                        {% for student in students %}
                            <li role='presentation'>
                                <a  href='#'
                                    style='cursor: pointer;'
                                    role='button'
                                    onClick='
                                        tt.ShowPointChart(
                                            this, "{{student.ID}}"
                                        )
                                    '>{{student.Name}}</a>
                            </li>
                        {% endfor %}
                    </ul>
                </div>
            </div>
        </div>
        <div class='{{col9}}'>
            <!-- Student point chart -->
            <div id='tt-pointchart' class='{{col12}} hidden'>
                <div class='panel panel-default'>
                    <div class='panel-heading'>
                        <span id="pointchart-heading"></span>
                        <span class="pull-right panel-subtitle"
                            >Student Overview</span>
                    </div>
                    <div class='panel-body'>
                        <canvas class='tt-panel-chart' class='{{col12}}'
                            ></canvas>
                        <!-- <div class='tt-panel-navbar'>
                            <button type='button' role='button'
                                class='btn btn-success tt-button-refresh'>

                                <span class='glyphicon glyphicon-refresh'
                                    ></span>
                            </button>
                        </div> -->
                    </div>
                </div>
            </div>
            <div id='tt-taskview' class='{{col12}} details-group hidden'>
                <div class='panel panel-default'>
                    <div class='panel-heading'>
                        <span id='taskview-heading'></span>
                        <span class='pull-right panel-subtitle'
                            >Task Result</span>
                    </div>
                    <div class='panel-body'>
                        <div class='panel'></div>
                    </div>
                </div>
            </div>
            <div id='tt-diffview' class='{{col12}} details-group hidden'>
                <div class='panel panel-default'>
                    <div class='panel-heading'>
                        <span id='diffview-heading'></span>
                        <span class='pull-right panel-subtitle'
                            >Result Diff</span>
                    </div>
                    <div class='panel-body'>
                        <div id='diffview-body' class='panel {{col10}}'></div>
                        <div id='diffview-menu'
                            class='panel {{col2}} pull-right'></div>
                    </div>
                </div>
            </div>
            <div id='tt-codeview' class='{{col12}} details-group hidden'>
                <div class='panel panel-default'>
                    <div class='panel-heading'>
                        <span id='codeview-heading'></span>
                        <span class='pull-right panel-subtitle'>Code View</span>
                    </div>
                    <div class='panel-body'>
                        <pre class="line-numbers">
                            <code class="language-clike"></code>
                        </pre>
                    </div>
                </div>
            </div>
        </div>
    </div>
{% endblock %}

{% block footer %}
    {% if view == "student" %}
        <footer class="footer">
            <div class="container">
                <button class="btn btn-primary tutortool-button"
                    onclick='window.location.href="/ui/grading"'>

                    <span class="glyphicon glyphicon-menu-left"
                        aria-hidden="true"></span>
                    <span class="text">Back to the Overview</span>
                </button>
                <!--
                <p class="text-muted">Place sticky footer content here.</p>
                -->
            </div>
        </footer>
    {% endif %}
{% endblock %}
