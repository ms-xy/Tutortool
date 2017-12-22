{% extends "base.tpl" %}

{% block additional_css %}
    <style>
        .detail-row>td {
            border-top: 1px solid #f5f5f5 !important;
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
    <script src="/static/js/views/tasks/main.js"></script>
{% endblock %}

{% block content %}
    <div class='row'>
        <div id='panel-tasks' class='panel panel-default'></div>
    </div>
{% endblock %}

{% block footer %}
{% endblock %}
