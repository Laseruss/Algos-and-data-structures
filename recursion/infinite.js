// No longer infinite Sadge
//

function countdown(i) {
  console.log(i);
  if (i <= 0) {
    return
  }
  countdown(i - 1);
}

countdown(100);
