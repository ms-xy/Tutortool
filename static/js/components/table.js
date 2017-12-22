class Table extends JQueryWrapper {
  constructor ($parent) {
    super();
    this.$el = $("<table>", {addClass: "table"});
    this.$tbody = $("<tbody>", {appendTo: this.$el});
    if ($parent) {
      this.$el.appendTo($parent);
    }
  }
  append (rows) {
    if (rows) {
      if (!$.isArray(rows)) {
        rows = [rows];
      }
      for (let row of rows) {
        this.$tbody.append(row.$el);
      }
    }
    return this;
  }
  appendTo ($parent) {
    this.$el.appendTo($parent);
    return this;
  }
  clear () {
    this.$el.children().remove();
    return this;
  }
}

class TableRow extends JQueryWrapper {
  constructor (datacells) {
    super();
    this.$el = $("<tr>");
    if (datacells) {
      this.append(datacells);
    }
  }
  append (datacells) {
    if (datacells) {
      if (!$.isArray(datacells)) {
        datacells = [datacells];
      }
      for (let datacell of datacells) {
        this.$el.append(datacell.$el);
      }
    }
    return this;
  }
  appendTo (table) {
    this.$el.appendTo(table.$el);
    return this;
  }
}

class TableDataCell extends JQueryWrapper {
  constructor (options) {
    super();
    this.$el = this._create$el();
    if (options) {
      if (options.text)     this.text(options.text);
      if (options.html)     this.html(options.html);
      if (options.colspan)  this.colspan(options.colspan);
      if (options.rowspan)  this.rowspan(options.rowspan);
      if (options.css)      this.css(options.css);
      if (options.addClass) this.addClass(options.addClass);
      if (options.cls)      this.addClass(options.cls); // alias for addClass
    }
  }
  _create$el () {
    return $("<td>");
  }
  append ($content) {
    this.$el.append($content);
    return this;
  }
  appendTo (tablerow) {
    this.$el.appendTo(tablerow.$el);
    return this;
  }
  colspan (num) {
    this.$el.attr("colspan", num);
    return this;
  }
  rowspan (num) {
    this.$el.attr("rowspan", num);
    return this;
  }
}

class TableHeaderCell extends TableDataCell {
  _create$el () {
    return $("<th>");
  }
}
