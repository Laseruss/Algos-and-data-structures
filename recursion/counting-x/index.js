// Count occurences of a given char in a string.

function countX(str, x='x') {
  
  if (str.length === 1) {
    if (str[0] === x) {
      return 1;
    }
    return 0;
  }

  if (str[0] === x) {
    return 1 + countX(str.slice(1,));
  }
  return countX(str.slice(1,));
}

console.log(countX('axbxcxd'));
