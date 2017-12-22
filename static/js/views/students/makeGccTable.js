/*
Helper function to create a table containing detail information about a given
gcc result.
*/
function makeGccTable (gccResult) {
  return (new Table()).addClass("result-table").append([
    new TableRow([
      new TableHeaderCell({
        text: "Gcc Result",
        css: {
          "width": 80,
          "white-space": "nowrap"
        }
      }),
      new TableHeaderCell()
    ]),
    new TableRow([
      new TableDataCell({text: "ExitCode"}),
      new TableDataCell({text: ""+gccResult.ExitCode})
    ]),
    new TableRow([
      new TableDataCell({text: "Error"}),
      new TableDataCell({text: (gccResult.Error) ? gccResult.Error : "-"})
    ]),
    new TableRow([
      new TableDataCell({text: "KillReason"}),
      new TableDataCell({
        text: (gccResult.KillReason) ? gccResult.KillReason : "-"
      })
    ]),
    new TableRow([
      new TableDataCell({text: "Stdout"}),
      (gccResult.Stdout) ?
        (new TableDataCell()).append(
          (new PreformattedText())
            .text(gccResult.Stdout.decodeBase64Unicode())
            .$el
        )
        :
        new TableDataCell({text: "-"})
    ]),
    new TableRow([
      new TableDataCell({text: "Stderr"}),
      new TableDataCell({
        html: (gccResult.Stderr) ?
          "<pre>"+gccResult.Stderr.decodeBase64Unicode()+"</pre>" : "-"
      })
    ])
  ]).$el
}
