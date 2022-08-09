// Calculate !n for a given n (n = 5  - 5 * 4 * 3 * 2 * 1)

function factorial(x) {
  if (x === 1) return 1;
  return x * factorial(x - 1);
}


console.log(factorial(10));
