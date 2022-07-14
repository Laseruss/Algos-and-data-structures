// sum an array of numbers using recursion.

function sum(arr) {
  // Base case, empty array.
  if (arr.length === 0) {
    return 0;
  }
  return arr[0] + sum(arr.slice(1));
}

console.log(sum([2, 5, 3, 2, 9]));
console.log(sum([]));
