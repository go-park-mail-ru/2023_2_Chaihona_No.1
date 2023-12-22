// Реализоватьсвойство `size`  у всех массивов,
// которое работало бы точно так же, как и `length`.
Object.defineProperty(Array.prototype, "size",
  {
    get: function () {
      return this.length;
    },
    set: function (newValue) {
      this.length = newValue;
    } 
  }
);

// #1
console.log([0, 1].size); // 2

// #2
var arr = [0, 1, 2];
arr.size = 0;
console.log(arr); // []
