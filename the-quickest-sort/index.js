// Implementation of quicksort.

function qSort(arr) {
  // Base case.
  if (arr.length < 2) {
    return arr;
  }

  const pivot = Math.floor(arr.length / 2)
  const less = arr.filter(n => n < arr[pivot]);
  const bigger = arr.filter(n => n > arr[pivot]);

  return qSort(less).concat(arr[pivot], qSort(bigger));
}


console.log(qSort([10, 5, 2, 3, 323, 2, 1, 45, -23, -231]));
