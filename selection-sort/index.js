// Implementation of selection sort algorithm
//
// Loop over the array or list, Add the smallest
// to new array, then the second smallest on next iteration
// etc..

function sortSmallToLarge(arr) {
  const sorted = [];

  // While there are elements left in the original array
  // we need to continue to loop through.
  // For each iteration we add the smallest num to the
  // sorted array.
  while (arr.length > 0) {
    let smallest = arr[0];
    let smallestIndex = 0;
    // Loop through the array to find the smallest number
    // and it's corresponding index.
    for (let i = 1; i < arr.length; i++) {
      if (smallest > arr[i]) {
        smallest = arr[i];
        smallestIndex = i;
      }
    }
    if (arr.length === 1) {
      sorted.push(arr[0]);
      break;
    }
    // Push the smallest value to the sorted array.
    sorted.push(smallest);

    // Slice the original array to remove the smallest value
    // for the next iteration of the loop.
    arr = arr.slice(0, smallestIndex).concat(arr.slice(smallestIndex + 1));
  }
  return sorted;
}

function selectionSortTwo(arr) {
  for (let i = 0; i < arr.length; i++) {
    let lowestInd = i;
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[lowestInd] > arr[j]) {
        lowestInd = j;
      }
    }
    if (lowestInd !== i) {
      let temp = arr[i];
      arr[i] = arr[lowestInd];
      arr[lowestInd] = temp;
    }
  }
  return arr;
}

const myArr = [1, 3, 8, 2, 34, 12, 143, 24];

console.log(sortSmallToLarge(myArr));
console.log(selectionSortTwo(myArr));
