class PreformattedText extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $("<pre>");
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
}
