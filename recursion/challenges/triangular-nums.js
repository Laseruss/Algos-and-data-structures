// Returns the triangular number for a given n.

function triangular(n) {
  if (n === 0) {
    return 0;
  }

  return n + triangular(n - 1);
}

console.log(triangular(15));
