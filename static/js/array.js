
Array.prototype.grep = function(fn) {
  let ls = [];
  for (let i=0; i<this.length; i++) {
    if (fn(this[i]) === true) {
      ls.push(this[i]);
    }
  }
  return ls;
};
Object.defineProperty(Array.prototype, "grep", {enumerable: false});

Array.prototype.grepFirst = function(fn) {
  let ls = [];
  for (let i=0; i<this.length; i++) {
    if (fn(this[i]) === true) {
      return this[i];
    }
  }
  return undefined;
};
Object.defineProperty(Array.prototype, "grepFirst", {enumerable: false});

Array.prototype.createMap = function(fn) {
  let obj = {}, key;
  for (let i=0; i<this.length; i++) {
    key = fn(this[i]);
    obj[key] = this[i];
  }
  return obj;
};
Object.defineProperty(Array.prototype, "createMap", {enumerable: false});

// alias for createMap as the name is ambiguous
Array.prototype.asObject = Array.prototype.createMap;
Object.defineProperty(Array.prototype, "asObject", {enumerable: false});

Array.prototype.asMap = function(fn) {
  let m = new Map(), val;
  for (let i=0; i<this.length; i++) {
    val = this[i];
    m.set(fn(val), val);
  }
  return m;
};
Object.defineProperty(Array.prototype, "asMap", {enumerable: false});
