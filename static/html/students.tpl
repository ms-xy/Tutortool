{% extends "base.tpl" %}

{% block additional_css %}
    <link href="/static/css/prism.min.css" rel="stylesheet">
    <style>
        table.table.task-info-table>tbody>tr>td {
          vertical-align: middle;
        }

        .task-info-name,
        .task-info-actions,
        .task-info-points,
        .task-info-comment {
          text-overflow: ellipsis;
          white-space: nowrap;
          overflow: hidden;
        }

        .task-info-name {
          width: 250px;
          max-width: 0px;
        }

        .task-info-actions {
          width: 150px;
          max-width: 0px;
        }

        .task-info-points {
          width: 70px;
          max-width: 0px;
        }

        .result-table {
            table-layout: fixed;
            overflow: auto;
            font-size: x-small;
        }

        .result-table pre {
            font-size: x-small;
            max-height: 400px;
            overflow: auto;
        }

        .source-file-listing {
            width: 100%;
            overflow: auto;
            font-size: x-small;
            resize: none;
        }

        .source-file-listing pre {
            font-size: x-small;
            max-height: 200px;
            overflow: auto;
        }

        .x-small-text {
            font-size: x-small;
        }

        #nav-student-wrapper {
            height: calc(100vh - 100px) !important;
            overflow: auto;
        }
    </style>
{% endblock %}

{% block additional_js %}
    <script src="/static/js/tutortool/api.js"></script>
    <script src="/static/js/tutortool/api-tasks.js"></script>
    <script src="/static/js/tutortool/api-commontasks.js"></script>
    <script src="/static/js/components/jquery.wrapper.js"></script>
    <script src="/static/js/components/button.js"></script>
    <script src="/static/js/components/viewport.js"></script>
    <script src="/static/js/components/table.js"></script>
    <script src="/static/js/components/row.js"></script>
    <script src="/static/js/components/col.js"></script>
    <script src="/static/js/components/panel.js"></script>
    <script src="/static/js/components/textarea.js"></script>
    <script src="/static/js/components/preformatted.js"></script>
    <script src="/static/js/views/students/tasks.js"></script>
    <script src="/static/js/views/students/makeGccTable.js"></script>
    <script src="/static/js/views/students/makeRunTable.js"></script>
    <script src="/static/js/views/students/detailview.js"></script>
    <!-- syntax highlightning -->
    <script type="text/javascript" src="/static/js/prism.min.js" data-manual></script>
    <!-- ajax error -->
    <script type="text/javascript">
        $(window).ready(() => {
            $.ajaxSetup({
                error: tutortool.showAjaxError
            });
        });
    </script>
    <script src="/static/js/views/students/main.js"></script>
{% endblock %}

{% block content %}
    <div class='row'>
        <div class='col-xs-3'>
            <div class='panel panel-default'>
                <div class='panel-body' id="nav-student-wrapper">
                    <ul id='nav-student' class='nav nav-pills nav-stacked'></ul>
                </div>
            </div>
        </div>
        <div class='col-xs-9'>
            <div id='panel-student' class='panel panel-default'></div>
            <!-- <div class='panel panel-primary'>
                <div id='detailview' class="panel-body"></div>
            </div> -->
            <div class='panel panel-warning'>
                <div id='panel-tasks' class="panel-body"></div>
            </div>
        </div>
    </div>
{% endblock %}

{% block footer %}
{% endblock %}
