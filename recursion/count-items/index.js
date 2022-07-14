// Count the items in a given list using recursion.
// Works for items even if there are multiple nested
// lists inside the given list.

function count(arr) {
  // Base case is an empty list.
  if (arr.length === 0) {
    return 0;
  }

  if (Array.isArray(arr[0])) {
    return count(arr[0]) + count(arr.slice(1));
  }

  return 1 + count(arr.slice(1));
}

console.log(count(["Hello", "World", "!", 
  ["I", "will", "be", "counted", "recursively"], 
  "Even", "nested", "array", "items", "are", "being", "counted"]));
