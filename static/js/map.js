
/*
create a new array based on the values of the map and sort these according to
the given comparator function
*/
Map.prototype.sort = function(compfn) {
  let arr = [];
  for (let entry of this) {
    arr.push(entry[1]); // [0] is the key, [1] is the value
  }
  return arr.sort(compfn);
};
Object.defineProperty(Map.prototype, "sort", {enumerable: false});
