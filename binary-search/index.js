// Takes a sorted array and an item that may
// or may not be in the array.

function binarySearch(arr, item) {
  let low = 0;
  let high = arr.length - 1;

  while (low <= high) {

    let mid = Math.floor((high + low) / 2);

    if (arr[mid] === item) {
      return `Value found at index: ${mid}`;
    } else if (arr[mid] < item) {
      low = mid + 1;
    } else if (arr[mid] > item) {
      high = mid - 1;
    }
  }
  return `${item} is not in the array`;
}

function recursiveBinarySearch(arr, item) {
  if (arr.length === 1) {
    return arr[0] === item;
  }

  if (arr[Math.floor(arr.length / 2)] === item) {
    return true;
  }

  return arr[Math.floor(arr.length / 2)] < item ?
    recursiveBinarySearch(arr.slice(Math.floor(arr.length / 2)), item) :
    recursiveBinarySearch(arr.slice(0, Math.floor(arr.length / 2)), item);
}

const myList = [1, 3, 5, 7, 9];

console.log(binarySearch(myList, 5));
console.log(recursiveBinarySearch(myList, 1));
