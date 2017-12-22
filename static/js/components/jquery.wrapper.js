class JQueryWrapper {
  addClass () {
    this.$el.addClass.apply(this.$el, arguments);
    return this;
  }
  removeClass () {
    this.$el.removeClass.apply(this.$el, arguments);
    return this;
  }

  append () {
    this.$el.append.apply(this.$el, arguments);
    return this;
  }

  appendTo () {
    this.$el.appendTo.apply(this.$el, arguments);
    return this;
  }

  css () {
    this.$el.css.apply(this.$el, arguments);
    return this;
  }

  attr () {
    this.$el.attr.apply(this.$el, arguments);
    return this;
  }

  show () {
    this.$el.show.apply(this.$el, arguments);
    return this;
  }
  hide () {
    this.$el.hide.apply(this.$el, arguments);
    return this;
  }

  html () {
    this.$el.html.apply(this.$el, arguments);
    return this;
  }
  text () {
    this.$el.text.apply(this.$el, arguments);
    return this;
  }

  click () {
    this.$el.click.apply(this.$el, arguments);
    return this;
  }
  on () {
    this.$el.on.apply(this.$el, arguments);
    return this;
  }
  off () {
    this.$el.off.apply(this.$el, arguments);
    return this;
  }

  children () {
    return this.$el.children.apply(this.$el, arguments);
    // NOT CHAINABLE!!!
  }

  clear () {
    this.$el.children().remove();
    return this;
  }

  remove () {
    this.$el.remove();
    // NOT CHAINABLE!!!
  }
}
