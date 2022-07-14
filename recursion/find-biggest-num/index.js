// Find the biggest number in a given array recursively.

function findBiggest(arr) {
  if (arr.length === 1) {
    return arr[0];
  }

  if (Array.isArray(arr[0])) {
    return findBiggest(arr[0]);
  }

  return Math.max(arr[0], findBiggest(arr.slice(1)));
}

console.log(findBiggest([3, 9,[3, 53, -1], 23, 1, 2]));
