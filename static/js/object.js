
Object.prototype.asMap = function(fn) {
  let m = new Map();
  if (fn) {
    for (let x of Object.entries(this)) {
      m.set(fn(x[1]), x[1]);
    }
  } else {
    for (let x of Object.entries(this)) {
      m.set(x[0], x[1]);
    }
  }
  return m;
};
Object.defineProperty(Object.prototype, "asMap", {enumerable: false});
