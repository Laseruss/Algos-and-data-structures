// Algorithm to find the intersection between two arrays in O(n).

function intersection(arr1, arr2) {
    let largestArr;
    let smallestArr;
    let hash = {};
    let intersection = [];

    if (arr1.length > arr2.length) {
        largestArr = arr1;
        smallestArr = arr2;
    } else {
        largestArr = arr2;
        smallestArr = arr1;
    }

    largestArr.forEach(i => {
        hash[i] = true;
    })

    smallestArr.forEach(i => {
        if (hash[i]) {
            intersection.push(i);
        }
    })

    return intersection;
}

console.log(intersection([1,2,3,4,5], [0,2,4,6,8]));
