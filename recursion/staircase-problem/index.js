// Number of possible paths up a stair of n steps if person can jump 1, 2 or 3 steps.

function numPaths(n) {
  if (n < 0) {
    return 0;
  }
  if (n === 1 || n === 0) {
    return 1;
  }
  return numPaths(n - 1) + numPaths(n - 2) + numPaths(n - 3);
}

console.log(numPaths(11));
