/*
Helper function to create a table containing detail information about a given
gcc result.
*/
function makeRunTable (testcase, runResult) { // TODO port to run result
  return (new Table()).addClass("result-table").append([
    new TableRow([
      new TableHeaderCell({
        text: "Testcase #"+testcase.Number+": "+testcase.Name,
        css: {
          "width": 80,
          "white-space": "nowrap"
        }
      }),
      new TableHeaderCell()
    ]),
    new TableRow([
      new TableDataCell({text: "ExitCode"}),
      new TableDataCell({text: ""+runResult.ExitCode})
    ]),
    new TableRow([
      new TableDataCell({text: "Error"}),
      new TableDataCell({text: (runResult.Error) ? runResult.Error : "-"})
    ]),
    new TableRow([
      new TableDataCell({text: "KillReason"}),
      new TableDataCell({
        text: (runResult.KillReason) ? runResult.KillReason : "-"
      })
    ]),
    new TableRow([
      new TableDataCell({text: "Stdout"}),
      (runResult.Stdout) ?
        (new TableDataCell()).append(
          (new PreformattedText())
            .text(runResult.Stdout.decodeBase64Unicode())
            // .on("scroll", function(ev) {
            //   var $this = $(this),
            //     scrollTop = this.scrollTop,
            //     scrollHeight = this.scrollHeight,
            //     height = $this.innerHeight(),
            //     delta = (ev.type == 'DOMMouseScroll' ?
            //       ev.originalEvent.detail * -40 :
            //       ev.originalEvent.wheelDelta),
            //     up = delta > 0;

            //   var prevent = function() {
            //     ev.stopPropagation();
            //     ev.preventDefault();
            //     ev.returnValue = false;
            //     return false;
            //   }

            //   if (!up && -delta > scrollHeight - height - scrollTop) {
            //     // Scrolling down, but this will take us past the bottom.
            //     $this.scrollTop(scrollHeight);
            //     return prevent();
            //   } else if (up && delta > scrollTop) {
            //     // Scrolling up, but this will take us past the top.
            //     $this.scrollTop(0);
            //     return prevent();
            //   }
            // })
            .$el
        )
        :
        new TableDataCell({text: "-"})
    ]),
    new TableRow([
      new TableDataCell({text: "Stderr"}),
      new TableDataCell({
        html: (runResult.Stderr) ?
          "<pre>"+runResult.Stderr.decodeBase64Unicode()+"</pre>" : "-"
      })
    ])
  ]).$el
}
