// Implementation of dijkstras algorithm.

/* const graph = {
    start: {
        a: 6,
	b: 2
	},
    a: {
	fin: 1
	},
    b: {
        a: 3,
	fin: 5
	},
    fin: {}
}

const costs = {
    a: 6,
    b: 2,
    fin: Infinity
}

const parents = {
    a: "start",
    b: "start",
    fin: null
}
*/

const graph = {
    start: {
	a: 5,
	b: 2,
    },
    a: {
	c: 4,
	d: 2
    },
    b: {
	a: 8,
	d: 7
    },
    c: {
	fin: 3,
	d: 6
    },
    d: {
	fin: 1
    },
    fin: {}
}

const costs = {
    a: 5,
    b: 2,
    c: Infinity,
    d: Infinity,
    fin: Infinity
}

const parents = {
    a: "start",
    b: "start"
}

function findLowestCostNode(costs) {
    let lowestCost = Infinity;
    let lowestCostNode = null;
    for (let [node, _] of Object.entries(costs)) {
	let cost = costs[node];
	if (cost < lowestCost && !processed.includes(node)) {
	    lowestCost = cost;
	    lowestCostNode = node;
	}
    }
    return lowestCostNode;
}

const processed = [];

let node = findLowestCostNode(costs);
while (node !== null) {
    let cost = costs[node];
    let neighbors = graph[node];
    for (let [key, _] of Object.entries(neighbors)) {
	let newCost = cost + neighbors[key]
	if (costs[key] > newCost) {
	    costs[key] = newCost;
	    parents[key] = node;
	}
    }
    processed.push(node);
    node = findLowestCostNode(costs);
}

function displayPath(obj) {
    const path = ["fin"];
    let x = obj["fin"];
    while (x !== "start") {
	path.unshift(x);
	x = parents[x];
    }
    path.unshift("start");
    return path;
}
console.log(displayPath(parents));
