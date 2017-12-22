class Textarea extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $("<textarea>");
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}
