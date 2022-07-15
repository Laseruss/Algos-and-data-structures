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
    this.table[index] = [key, value];
    this.size++;
  }

  get(key) {
    return this.table[this._generateHash(key)];
  }

  remove(key) {
    const index = this._generateHash(key);
    this.table[index] = undefined;
  }
}


const table = new HashTable(127);

table.set("orange", 2);
table.set("banana", 10);
table.set("apple", 4);

console.log(table.get("banana"));
console.log(table.get("apple"));

table.remove("apple");

console.log(table.get("apple"));
console.log(table.get("starfruit"));
