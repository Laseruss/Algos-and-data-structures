// Function to count down to zero from a given integer.

function countdown(num) {
  if (num === 0) {
    return "BlastOff!!!!!";
  }

  console.log(num)
  return countdown(num - 1);
}

console.log(countdown(10));
