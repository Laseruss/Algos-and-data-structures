function bubbleSort(arr) {
    let maxIndex = arr.length - 1;
    let sorted = false;

    while (true) {
        sorted = true;

        for (let i = 0; i < maxIndex; i++) {
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
        maxIndex--;
    }
}

const arr = [65, 44, 45, 35, 25, 15, 10];

console.log(bubbleSort(arr));
