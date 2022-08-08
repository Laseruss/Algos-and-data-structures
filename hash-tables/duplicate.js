// Function to find the first duplicate char in a given array of strings.

function duplicate(arr) {
    const hash = {};

    for (let i = 0; i < arr.length; i++) {
        if (hash[arr[i]]) {
            return arr[i];
        }

        hash[arr[i]] = true;
    }

    return "No duplicate strings in the given array.";
}

console.log(duplicate(["a", "b", "c", "d", "c", "e"]));
console.log(duplicate(["a", "b", "d", "c", "e"]));
