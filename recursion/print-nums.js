// Print numbers of nested arrays.


function printNums(arr) {
  if(arr.length === 0) {
    return;
  }

  if(Array.isArray(arr[0])) {
    return printNums(arr[0]) + printNums(arr.slice(1));
  }
  console.log(arr[0]);
  return printNums(arr.slice(1));
}
const array = [
  1,
  2,
  3,
  [4, 5, 6],
  7,
  [8,
    [9, 10, 11,
      [12, 13, 14]
    ]
  ],
  [15, 16, 17, 18, 19,
    [20, 21, 22,
      [23, 24, 25,
        [26, 27, 29]
      ], 30, 31
    ], 32 
  ], 33
]
printNums(array);
