class Button {

  constructor ($parent) {
    this.$el = $("<button>")
      .addClass("btn btn-default")
      .data("tutortool.button", this);

    this._type = "btn-default";
    this._tooltipOptions = {
      title: undefined,
      container: "body",
      toggle: "tooltip",
      placement: "auto top",
      html: true,
    };

    if ($parent) {
      this.$el.appendTo($parent);
    }
  }

  appendTo ($parent) {
    this.$el.appendTo($parent);
    return this;
  }

  addClass (cls) {
    this.$el.addClass(cls);
    return this;
  }

  removeClass (cls) {
    this.$el.removeClass(cls);
  }

  /*
  Set (or unset) the glyphicon to use for the button. The glyphicon will always
  be located on the left of any text inside of the button.
  */
  glyphicon (glyphicon) {
    // remove old glyphicon
    if (this.$glyphicon) {
      this.$glyphicon.remove();
      this.$glyphicon = undefined;
    }

    // create new and add
    if (glyphicon) {
      this.$glyphicon = $("<span>").addClass("glyphicon glyphicon-"+glyphicon);
      this.$el.prepend(this.$glyphicon);
    }

    return this;
  }

  /*
  Set a text to be displayed inside of the button. It will always be on the
  right of any glyphicon present.
  */
  text (text) {
    // remove old text
    if (this.$text) {
      this.$text.remove();
      this.$text = undefined;
    }

    // create new and add
    if (text) {
      this.$text = $("<span>").html(text);
      this.$el.append(this.$text)
    }

    return this;
  }

  /*
  Set the button type. This only changes the buttons color.
  Applicable types are:
  - default
  - primary
  - info
  - success
  - warning
  - danger
  */
  type (newtype) {
    this.$el.removeClass(this._type);
    this._type = "btn-"+newtype;
    this.$el.addClass(this._type);
    return this;
  }

  /*
  Set or unset the tooltip of the button. Requires some "hacks" to make it work.
  */
  tooltip (text) {
    if (!text && this._tooltipOptions.title) {
      this._tooltipOptions.title = undefined;
      this.$el.tooltip('destroy');
    }
    if (text) {
      if (!this._tooltipOptions.title) {
        this._tooltipOptions.title = text;
        this.$el.tooltip(this._tooltipOptions);
      } else {
        this.$el.attr("title", text);
        this.$el.tooltip("fixTitle");
        // this.$el.tooltip("show");
      }
    }
    return this;
  }

  /*
  Set / remove the event-click handler. It is a glorified wrapper for the
  underlying JQuery click api.
  */
  click (fn) {
    if (!fn) {
      this.$el.off("click")
    } else {
      this.$el.click(fn);
    }
    return this;
  }

  /*
  Change the button background. Useful for displaying a progress bar for
  compilation and run buttons.
  */
  background (picturepath) {
    if (picturepath) {
      this.$el.css("background-image", 'url("'+picturepath+'")');
    } else {
      this.$el.css("background-image", "none");
    }
    return this;
  }
}
