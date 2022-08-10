// Function that sums the characters in an array of strings.

function sumChars(arr) {
  if (arr.length === 1) {
    return arr[0].length;
  }

  return arr[0].length + sumChars(arr.slice(1));
}

console.log(sumChars(['ab', 'c', 'def', 'ghij']));
