// Implementation of dijkstras algorithm.

const graph = {
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
    console.log(node);
    node = findLowestCostNode(costs);
}
