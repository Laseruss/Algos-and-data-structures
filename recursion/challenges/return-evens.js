// Function that returns the even numbers of a given array of integers.

function onlyEven(arr) {
  // Base case
  if (arr.length === 0) {
    return [];
  }
  
  if (arr[0] % 2 === 0) {
    return [arr[0]] + onlyEven(arr.slice(1));
  }

  // Call onlyEven on rest of array
  return onlyEven(arr.slice(1));
}

console.log(onlyEven([1,2,3,4,5,6]));
