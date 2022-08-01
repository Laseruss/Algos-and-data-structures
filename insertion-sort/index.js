function insertionSort(arr) {
    for (let i = 1; i < arr.length; i++) {
        let temp = arr[i];
        let pos = i - 1;

        while (pos >= 0) {
            if (arr[pos] > temp) {
                arr[pos + 1] = arr[pos];
                pos--;
                continue;
            }
            break;
        }
        arr[pos + 1] = temp;
        console.log(arr);
    }
    return arr;
}

const nums = [4, 2, 7, 1, 3];

console.log(insertionSort(nums));
