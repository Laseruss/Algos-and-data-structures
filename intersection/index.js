const arr1 = [1, 6, 34, 0, 8, 3];
const arr2 = [4, 5, 1, 54, 89, 3];

function findIntersect(arr1, arr2) {
    const res = [];
    for (let i = 0; i < arr1.length; i++) {
        for (let j = 0; j < arr2.length; j++) {
            if (arr1[i] === arr2[j]) {
                res.push(arr1[i]);
            }
        }
    }
    return res;
}

// Faster implementation in the average and best case.
// By not continuing to iterate over arr2 if we find a match.
function findIntersectFaster(arr1, arr2) {
    const res = [];
    for (let i = 0; i < arr1.length; i++) {
        for (let j = 0; j < arr2.length; j++) {
            if (arr1[i] === arr2[j]) {
                res.push(arr1[i]);
                break;
            }
        }
    }
    return res;
}

console.log(findIntersect(arr1, arr2));
console.log(findIntersectFaster(arr1, arr2));
