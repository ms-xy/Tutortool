class Viewport {
  constructor (zIndex, $children) {
    /*
    Create the viewport element.
    It is immediately attached to the body element of the DOM, but hidden.
    */
    this.$el = $("<div>", {
      "appendTo": $("body"),
      "css": {
        "position": "fixed",
        "top":    0,
        "right":  0,
        "bottom": 0,
        "left":   0,
        "background-color": "white",
        "z-index": zIndex, // should be well above bootstraps indices
        "overflow": "auto"
      }
    }).hide();

    /*
    Create a function for removing the viewport by pressing escape.
    This function is proxied to this. Alternatively could set a local variable,
    but this is more canonical with the use of jquery on the rest of the site.
    */
    this._hide_by_escape = $.proxy(function(ev) {
      if (ev.key == "Escape") {
        this.hide();
      }
    }, this)


    /*
    Event handling
    */
    this._events = {};

    /*
    If children are passed via constructor, append them right away
    */
    if ($children) {
      this.$el.append($children);
    }
  }

  on (event, fn) {
    if (!this._events[event]) {
      this._events[event] = new Map();
    }
    if (!this._events[event].get(fn)) {
      this._events[event].set(fn, true);
    }
    return this;
  }
  off (event, fn) {
    if (this._events[event]) {
      if (fn) {
        this._events[event].delete(fn);
      } else {
        delete this._events[event];
      }
    }
    return this;
  }
  trigger (event, data) {
    if (this._events[event]) {
      for (let entry of this._events[event]) {
        let fn = entry[0]; // fn is the key not the value
        fn(data);
      }
    }
    return this;
  }

  show () {
    this.$el.show();
    $(window).on("keydown", window, this._hide_by_escape);
    this.trigger("show");
    return this;
  }

  hide () {
    this.$el.hide();
    $(window).off("keydown", window, this._hide_by_escape);
    this.trigger("hide");
    return this;
  }

  clear () {
    this.$el.children().remove();
    this.trigger("clear");
    return this;
  }

  destroy () {
    this.$el.remove();
    this.trigger("destroy");
    return this;
  }

  append ($content) {
    this.$el.append($content);
    this.trigger("append");
    return this;
  }
}
