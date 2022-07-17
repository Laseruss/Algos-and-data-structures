// Implement the graph.
// I think this is an implementation of bfs

graph = {
  you: ["alice", "bob", "claire"],
  alice: ["peggy"],
  bob: ["anuj", "peggy"],
  claire: ["thom", "jonny"],
  anuj: [],
  peggy: [],
  thom: [],
  jonny: []
}

function endsWithM(x) {
  return x.endsWith("m");
}

let queue = graph["you"];
let searched = [];
const wanted = "thom";

while (queue.length !== 0) {
  if (searched.includes(queue[0])) {
    queue.shift();
    continue;
  }
  if (endsWithM(queue[0])) {
    console.log(queue.shift());
    break;
  }
  searched.push(queue[0]);
  queue = queue.concat(graph[queue[0]]);
  queue.shift();
}
