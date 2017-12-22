
class Row extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $('<div class="row"></div>');
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}
