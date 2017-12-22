
class Column extends JQueryWrapper {
  constructor (width, $parent) {
    super();
    let _width = width || 12;
    this.$el = $('<div class="col-xs-'+_width+'"></div>');
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}
