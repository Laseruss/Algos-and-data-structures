// Reverse a string using a stack.

class Stack {
    constructor () {
        this.stack = [];
        this.len = 0;
    }

    push(val) {
        this.len++;
        this.stack.push(val);
    }

    pop() {
        this.len--;
        return this.stack.pop();
    }
}

function reverse(str) {
    let stack = new Stack();
    const strArr = str.split("");

    strArr.forEach(c => stack.push(c));

    let res = "";
    while (stack.len !== 0) {
        res += stack.pop();
    }
    return res;
}

console.log(reverse("hello my name is olleh"));
