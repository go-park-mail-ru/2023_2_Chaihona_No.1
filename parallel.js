// Параллельные вычисления

function parallel(arr, finish) {
  if (!Array.isArray(arr)) {
      throw new Error('wrong arguments');
  }

  if (typeof finish !== 'function') {
      throw new Error('wrong arguments');
  }

  let results = [];
  let jobs = arr.length;

  if (jobs === 0) {
      finish(results);
      return;
  }

  arr.forEach(function (fn, i) {
      if (typeof fn !== 'function') {
          throw new Error('wrong arguments');
      }
      let res = fn((value) => {
        results[i] = value;
        jobs -= 1;
        if (jobs === 0) {
          finish(results);
        }
      });
      if (res !== undefined) {
        jobs -= 1;
        results[i] = res;
        if (jobs === 0) {
          finish(results);
        }
      }
  });
}


parallel([
function (resolve) {
  setTimeout(function () {
    resolve(10);
  }, 5000);
},
function () {
  return 5;
},
function (resolve) {
  setTimeout(function () {
    resolve(0);
  }, 1000)
}
], function (results) {
console.log(results); // [10, 5, 0]
// console.log(results.length);
});
