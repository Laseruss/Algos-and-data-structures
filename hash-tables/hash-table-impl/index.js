// Making my own hash table generator.

class LinkedList{
  constructor() {
    this.head = null;
  }

  addNode(data) {
    const newNode = new Node(data, this.head);
    this.head = newNode;
  }

  findNode(key) {
    let curr = this.head;
    while (curr !== null) {
      if (curr.data[0] === key) {
        return curr.data;
      }
      curr = curr.next;
    }
  }
}

class Node{
  constructor(data, next = null) {
    this.data = data;
    this.next = next;
  }
}

class HashTable {
  constructor(len) {
    this.table = new Array(len);
    this.size = 0;
  }

  _generateHash(key) {
    return key.split("").reduce((a,b) => a + b.charCodeAt(), 0) % this.table.length;
  }

  set(key, value) {
    const index = this._generateHash(key);
    if (typeof this.table[index] === "undefined") {
      const list = new LinkedList();
      list.addNode([key, value]);
      this.table[index] = list;
      this.size++
    } else {
      this.table[index].addNode([key, value]);
      this.size++;
    }
  }

  setMany(arr) {                // Arr needs to be in form [[key, value], [key, value] etc..].
    arr.forEach(item => {
      this.set(item[0], item[1]);
      this.size++;
    })
  }
  // TODO: Sort of works, cant seem to make it so that i can iterate over the linked list multiple times.
  // Fix for another day
  get(key) {
    const index = this._generateHash(key);
    if (!this.table[index]) {
      return `key: ${key} is undefined`;
    }

    return this.table[index].findNode(key);
  }

  remove(key) {
    const index = this._generateHash(key);
    this.table[index] = undefined;
  }
}

const table = new HashTable(15);

table.set("orange", 2);
table.set("banana", 10);
table.set("apple", 4);
console.log(table.get("banana"));

table.remove("apple");
console.log(table.get("apple"));

table.setMany([["pear", 2], ["avocado", 29], ["onion", 1]])
console.log(table.get("avocado"));


table.set("Spain", 3);
table.set("ǻ", 5);
console.log(table.get("Spain"));
console.log(table.get("ǻ"));
