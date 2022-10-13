// Function that takes an array, and doubles each value in place.

function doubleInPlace(arr, i=0) {
  if (i >= arr.length) {
    return arr;
  }

  arr[i] = arr[i] * 2;
  return doubleInPlace(arr, i + 1);
}
const arr = doubleInPlace([1,2,3,4])
console.log(arr);
