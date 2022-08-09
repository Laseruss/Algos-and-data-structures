// Implementation of a queue using an underlying array.

class Queue {
    constructor(arr = []) {
        this.queue = arr;
    }

    enqueue(val) {
        this.queue.push(val);
    }

    dequeue() {
        return this.queue.shift();
    }

    read() {
        return this.queue[0];
    }
}

let queue = new Queue();

queue.enqueue(1);
queue.enqueue(2);
queue.enqueue(3);
queue.enqueue(4);
queue.enqueue(5);

let deq = queue.dequeue();

console.log(deq, queue.read());
