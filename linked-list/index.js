// First attempt at building a linked list
//

class Node{
  constructor (data, next = null){
    this.data = data;
    this.next = next;
  }
}

class LinkedList{
  constructor() {
    this.head = null;
  }

  insertAtBeginning(data) {
    // A newNode object is created with property data and next = null
    let newNode = new Node(data);

    // The pointer next is assigned head pointer so that both pointers
    // now point at the same node.
    newNode.next = this.head

    // As we are inserting at the beginning the head pointer needs to now
    // point at the new node.
    
    this.head = newNode;

    return this.head;
  }

  insertAtEnd(data) {
    // A new node object is created with property data and next = null.

    let newNode = new Node(data);

    // When head = null i.e. the list is empty, then head itself will point
    // to the new node.

    if (!this.head) {
      this.head = newNode;
      return this.head;
    }

    // Else, traverse the list to find the tail (the tail node will initially
    // be pointing at null), and update the tail's next pointer.
    let tail = this.head;
    while(tail.next !== null) {
      tail = tail.next;
    }
    tail.next = newNode;

    return this.head;
  }

  find(i) {
    let counter = 0;
    let curr = this.head;
    while (this.head !== null) {
      if (counter === i) {
        return curr;
      }
      counter++;
      curr = curr.next;
    }
  }

  insert(data, i) {
    let nodeBefore = this.find(i);
    let after = nodeBefore.next;

    let newNode = new Node(data, after);
    nodeBefore.next = newNode;
  }


}

let list = new LinkedList;

const arr = [1,2,3,4,5,6,7,9];

arr.forEach(n => {
  list.insertAtBeginning(n);
})

list.insert(8, 3);

while (list.head !== null) {
  console.log(list.head);
  list.head = list.head.next;
}



