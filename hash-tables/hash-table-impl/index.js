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
    console.log(index);
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
      const index = this._generateHash(item[0]);
      this.table[index] = [item[0], item[1]]; 
      this.size++;
    })
  }

  get(key) {
    const index = this._generateHash(key);
    console.log(this.table[index]);
    if (!this.table[index]) {
      return undefined
    }
    if (this.table[index][0][0] === key) {
      return this.table[index];
    }
    for (let i = 0; i < this.table[index].length; i++) {
      if (this.table[index][i][0] === key) {
        return this.table[index][i];
      }
    } 
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

table.setMany([["pear", 2], ["avocado", 29], ["onion", 1]])
console.log(table.get("avocado"));


table.set("Spain", 3);
table.set("ǻ", 5);

console.log(table.get("Spain"));
console.log(table.get("ǻ"));
