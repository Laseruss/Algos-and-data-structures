// Takes a sorted array and an item that may
// or may not be in the array.
  // Binary search or always pick the middle
  // element until you find what you are
  // looking for. 
  // Runs in O(log n) time, unlike simple search
  // that walks through every element and runs
  // in O(n).
  // Big O notation shows the worst possible outcome.
function binarySearch(arr, item) {
  // Set the low starting point to the first
  // element of the array.
  let low = 0;
  // Set the high starting point to the last
  // element of the array.
  let high = arr.length - 1;

  // While the low element is smaller or equal
  // to the requested item, we havent searched
  // through the given array.
  while (low <= high) {
    // The current mid point of the elements
    // we still need to check.
    let mid = Math.floor((high + low) / 2);
    // If we found the item, return the index
    // where it was found.
      // If the item is larger or smaller then
      // the currently selected element we
      // need to change the search boundaries (low and mid).
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


const myList = [1, 3, 5, 7, 9];

console.log(binarySearch(myList, 2));
