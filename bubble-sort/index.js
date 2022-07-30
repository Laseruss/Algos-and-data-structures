function bubbleSort(arr) {
    while (true) {
        let sorted = true;
        for (let i = 0; i < arr.length - 1; i++) {
             if(arr[i] > arr[i + 1]) {
                const temp = arr[i];
                arr[i] = arr[i + 1];
                arr[i + 1] = temp
                sorted = false;
            }
        }
        if (sorted) {
            return arr;
        }
    }
}

const arr = [-1234, 34, 1, 577, 453, 1, 5, 22, 985, 21, 2, 34, 6, 3, 5];

console.log(bubbleSort(arr));
