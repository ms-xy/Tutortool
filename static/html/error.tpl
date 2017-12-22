{% extends "base.tpl" %}

{% block additional_css %}
    <!-- <link href="static/css/tutortool_studentdetails.css" rel="stylesheet"> -->
{% endblock %}

{% block additional_js %}
    <!-- <script src="static/js/tutortool_studentdetails.js"></script> -->
{% endblock %}

{% block content %}
    {% set col3="tutortool-panel col-sm-3 col-md-3 col-lg-3" %}
    {% set col6="tutortool-panel col-sm-6 col-md-6 col-lg-6" %}
    {% set col9="tutortool-panel col-sm-9 col-md-9 col-lg-9" %}
    {% set col12="tutortool-panel col-sm-12 col-md-12 col-lg-12" %}

    <div class='row'>
        <div class='{{col12}}'>
            <div class='panel panel-danger'>
                <div class='panel-heading'>
                    {{error_code}}
                </div>
                <div class='panel-body text-danger'>
                    {{error_msg}}
                </div>
            </div>
        </div>
    </div>
{% endblock %}

{% block footer %}
    {% if view == "student" %}
        <footer class="footer">
            <div class="container">
                <button class="btn btn-primary tutortool-button" onclick='window.location.href="/ui/grading"'>
                    <span class="glyphicon glyphicon-menu-left" aria-hidden="true"></span>
                    <span class="text">Back to the Overview</span>
                </button>
                <!-- <p class="text-muted">Place sticky footer content here.</p> -->
            </div>
        </footer>
    {% endif %}
{% endblock %}
