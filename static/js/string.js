
String.prototype.decodeBase64Unicode = function () {
  return decodeURIComponent(Array.prototype.map.call(atob(this), function(c) {
    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
  }).join(''))
}
Object.defineProperty(String.prototype, "decodeBase64Unicode", {enumerable: false});

String.prototype.encodeBase64Unicode = function () {
    return btoa(encodeURIComponent(this).replace(/%([0-9A-F]{2})/g, function(match, p1) {
        return String.fromCharCode(parseInt(p1, 16))
    }))
}
Object.defineProperty(String.prototype, "encodeBase64Unicode", {enumerable: false});
