// Find the greatest common denominator between two positive integers usign
// the Euclidean algorithm.

function euclid(a, b) {

  if (a === 0 || b === 0) {
    return a === 0 ? b : a;
  }

  const remainder = a > b ? a % b : b % a;
  return euclid(a > b ? b: a, remainder);
}

console.log(euclid(1680, 640));
