// Making my own hash table generator.

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
      this.table[index] = [[key, value]];
      this.size++
    } else {
      this.table[index].push([key,value]);
      this.size++;
    }
  }

  setMany(arr) {                // Arr needs to be in form [[key, value], [key, value] etc..].
    arr.forEach(item => {
      this.set(item[0], item[1]);
      this.size++;
    })
  }

  get(key) {
    const index = this._generateHash(key);
    if (!this.table[index]) {
      return `key: ${key} is undefined`;
    }
    if (this.table[index][0][0] === key && this.table[index].length === 1) {
      return this.table[index];
    }
    for (let i = 0; i < this.table[index].length; i++) {
      if (this.table[index][i][0] === key) {
        return this.table[index][i];
      }
    } 
  }

  // Remove does not work if there has been collisions on that index of the array
  // will remove all elements at that index and set value to undefined.
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
