<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
        <meta name="description" content="">
        <meta name="author" content="Maximilian Schott, tutortool@roottec.com">
        <link rel="icon" href="../../favicon.ico">

        <title>Tutortool</title>

        <!-- Bootstrap core CSS -->
        <link href="/static/bootstrap/css/bootstrap.min.css" rel="stylesheet">

        <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
        <link href="/static/css/ie10-viewport-bug-workaround.css" rel="stylesheet">

        <!-- Custom styles for this template -->
        <!-- <link href="/static/css/sticky-footer-navbar.css" rel="stylesheet"> -->
        <link href="/static/css/tutortool.css" rel="stylesheet">
        {% block additional_css %}{% endblock %}

        <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
        <!--[if lt IE 9]>
            <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
            <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
        <![endif]-->
        <style>
            body > .container {
              padding: 60px 15px 0;
            }
        </style>
    </head>

    <body>

        <!-- Error view (floating, centered) -->
        <div id='tt-error' class='{{col6}} hidden'>
            <div class='panel panel-danger floating-centered'>
                <div class='panel-heading'>
                    <span></span>
                    <button role='button' class='btn btn-danger pull-right btm-sm' onclick='$("#tt-error").addClass("hidden");' style='margin-top: -7px'>
                        <span class='glyphicon glyphicon-remove'></span>
                    </button>
                </div>
                <div class='panel-body'></div>
            </div>
        </div>

        <!-- Fixed navbar -->
        <nav class="navbar navbar-default navbar-fixed-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="#">Tutortool</a>
                </div>
                <div id="navbar" class="collapse navbar-collapse">
                    <ul class="nav navbar-nav">

                        <li class='{{view_index_active}}'>
                            <a href="/ui/index">Home</a></li>

                        <li class='{{view_students_active}}'>
                            <a href="/ui/students">Students</a></li>

                        <li class='{{view_tasks_active}}'>
                            <a href="/ui/tasks">Tasks</a></li>

                        <!-- <li class="dropdown">
                            <a class='{{ settings|default:"" }}' href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Settings <span class="caret"></span></a>
                            <ul class="dropdown-menu">
                                <li><a href="#">Action</a></li>
                                <li><a href="#">Another action</a></li>
                                <li><a href="#">Something else here</a></li>
                                <li role="separator" class="divider"></li>
                                <li class="dropdown-header">Nav header</li>
                                <li><a href="#">Separated link</a></li>
                                <li><a href="#">One more separated link</a></li>
                            </ul>
                        </li> -->
                    </ul>
                </div><!--/.nav-collapse -->
            </div>
        </nav>

        <!-- Begin page content -->
        <div class="container">
            {% block content %}{% endblock %}
        </div>

        <!-- Optional page footer -->
        {% block footer %}{% endblock %}

        <!-- Bootstrap core JavaScript
        ================================================== -->
        <!-- Placed at the end of the document so the pages load faster -->
        <script src="/static/css/jquery-3.1.1.min.js"></script>
        <script src="/static/bootstrap/js/bootstrap.min.js"></script>
        <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
        <script src="/static/js/ie10-viewport-bug-workaround.js"></script>
        <!-- Chart.bundle.min.js contains Chart.js and Moment.js -->
        <script src="/static/js/Chart.bundle.min.js"></script>
        <script src="/static/js/array.js"></script>
        <script src="/static/js/object.js"></script>
        <script src="/static/js/map.js"></script>
        <script src="/static/js/string.js"></script>
        <script src="/static/js/tutortool/tutortool.js"></script>
        {% block additional_js %}{% endblock %}
    </body>
</html>
