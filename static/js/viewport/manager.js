function Viewport ($el) {
  // clean viewport
  $el.empty(true);

  // keep reference
  this.$el = $el;

  this.render = function(document_fragment) {
    this.$el.empty().append(document_fragment)
  }
}
