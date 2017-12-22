
class Panel extends JQueryWrapper {
  constructor (type, $parent) {
    super();
    let _type = type || "default";
    this.$el = $('<div class="panel panel-'+_type+'"></div>');
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }

  get heading () {
    if (!this._heading) {
      this._heading = new PanelHeading(this.$el);
    }
    return this._heading;
  }

  get body () {
    if (!this._body) {
      this._body = new PanelBody(this.$el);
    }
    return this._body;
  }
}

class PanelHeading extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $('<div class="panel-heading"></div>');
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}

class PanelBody extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $('<div class="panel-body"></div>');
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}
