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

    read() {
        return this.stack[this.stack.length - 1];
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
            stack.push(arr[i]);
        } else if (pairs[arr[i]]) {
            val = stack.pop();
            if (!val) {
                return `${arr[i]} does not have a matching opening bracket`;
            } else if (pairs[arr[i]] !== val) {
                return `${arr[i]} has mismatched opening bracket`;
            }
        }
    }
    if (stack.read() !== undefined) {
        return `${stack.pop()} does not have a matching closing bracket.`;
    }

    return true;
}

console.log(match("[({})]"));
