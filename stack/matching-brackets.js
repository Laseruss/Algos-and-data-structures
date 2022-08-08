// Linter to show if a line has the correct opening and closing brackets (([{}])).

class Stack {
    constructor(arr = []) {
        this.stack = arr;
    }

    push(val) {
        this.stack.push(val);
    }

    pop() {
        return this.stack.pop();
    }
}

const pairs = {
    ")": "(",
    "]": "[",
    "}": "{"
}

function match(str) {
    const arr = str.split("");
    const stack = new Stack();

    for (let i = 0; i < arr.length; i++) {
        let p = "([{";

        if (p.includes(arr[i])) {
            stack.push(arr[i])
        } else if (pairs[arr[i]]) {
            if (pairs[arr[i]] !== stack.pop()) {
                return false;
            }
        }
    }
    return true;
}

console.log(match("(let[ x p) {asl} eggg [({})]"));
