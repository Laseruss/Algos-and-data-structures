// Implementation of a stack, using an underlying array. 

class Stack{ 
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
    ;}
}

let stack = new Stack();

stack.push(5);

stack.push("hello");

console.log(stack.pop());

console.log(stack.read());
