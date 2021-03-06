## Description

This is a grading tool for C(98/99/11) assignments. It provides a
Web-GUI which is provided by a simple integrated webserver.
The webserver is currently hard-coded to listen to port 8080.

## Usage

1. Open a terminal and run:

    ```sh
    go get "github.com/ms-xy/Tutortool"
    go get "github.com/ms-xy/rlimiter"
    ```
   
1. Switch to the directory you want to install the Tutortool in
1. Create the configuration (see config section below)
1. Copy the `rlimiter` and `Tutortool` executable to this directory:
    
    ```sh
    cp "$GOPATH/bin/Tutortool" "$GOPATH/bin/rlimiter" ./
    ```
    This step might seem unnecessary.
    However, it protects against accidentally overwriting
    the binaries in the `GOPATH`.
    
1. Execute the Tutortool:
    
    ```sh
    ./Tutortool
    ```

## Configuration

The Tutortool is configured via multiple JSON files.
Mandatory files are the `tutortool-config.json`, the students config and the
task list config.

### tutortool-config.json

```json
{
  "tutor":           "<name>",
  "students-config": "students/students-config.json",
  "task-list":       "tasks/task-list.json"
}
```

The `students-config` and `task-list` can be located anywhere. You can use
both relative and absolute paths here.

### students-config.json

```json
[
  {
    "name": "Anne",
    "path": "anne"
  },
  {
    "name": "Herbert",
    "path": "herbert"
  },
  {
    "name": "Lucas",
    "path": "lucas"
  }
]
```

This file is a simple map of directories. Relative paths are resolved in
relation to the path of the `students-config.json`.

### task-list.json

```json
[
  "task01/task-config.json",
  "task02/task-config.json",
  "task03/task-config.json"
]
```

This file lists all the task config files. Again, relative paths are resolved
relative to the `task-list-json` path.

### Example task-config.json

```json
{
  "name": "task name / short description",
  "gcc": {
    "parameters": ["-Wall", "-Werror", "-std=c99"],
    "files":      ["main.c", "library.c"],
    "replacements": {
      "main.c":    "main.c",
      "library.h": "library.h"
    }
  },
  "run": {
    "timeout": 60,
    "stdout-size": 10000,
    "stderr-size": 10000
  },
  "testcases": [
    {
      "name": "testcase name / short description",
      "parameters": ["custom", "program", "parameters"],
      "input-file": "path/to/input/file"
    },
    {
      "name": "another testcase",
      "input-file": "path/to/another/input/file",
      "timeout": 120,
      "stdout-size": 1000,
      "stderr-size": 2000
    },
    {
      "name": "just simply execute the darn program"
    }
  ],
  "sources": ["library.c"]
}
```

Each config file must specify a task name, gcc settings and at least one
testcase.
A global run config may be specified but also overriden on a per-testcase level.
Sane defaults are chosen if limits are not set. (5 min execution time, 10kB
stdout and stderr sizes, memory limit of a couple megabyte).
These limits are enforced by using [rlimiter](https://github.com/ms-xy/rlimiter).

## License

Tutortool itself is licensed under GNU GPLv3.
Please see the attached License.txt file for details.
Different license terms can be arranged on request.

Tutortool comes packaged with several javascript libraries that use different
licenses. These libraries remain subject to the terms and conditions set by
their authors and/or publishers. The user of the Tutortool acknowledges that
they have read and agreed to the terms and conditions of these libraries:

- [bootstrap](https://github.com/twbs/bootstrap/blob/master/LICENSE)
- [jquery](https://jquery.org/license)
- [prism.js](https://github.com/PrismJS/prism/blob/gh-pages/LICENSE)
- [Chart.js](https://github.com/chartjs/Chart.js/blob/master/LICENSE.md)
- diff.js (BSD License, see file [static/js/diff.min.js](static/js/diff.min.js))
- IE10 viewport adjustments by Twitter ([https://github.com/twbs/bootstrap/blob/master/LICENSE](https://github.com/twbs/bootstrap/blob/master/LICENSE))
