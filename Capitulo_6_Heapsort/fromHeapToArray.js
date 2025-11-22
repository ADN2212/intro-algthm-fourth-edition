//Esta es una lista de adjacencia que representa el arbol binario (heap):
const heapTree = [
	{
		val: 90,
		vecinos: [1, 2]
	},
	{
		val: 89,
		vecinos: [3, 4]
	},
	{
		val: 70,
		vecinos: [5, 6]
	},
	{
		val: 36,
		vecinos: [7, 8]
	},
	{
		val: 75,
		vecinos: []
	},
	{
		val: 63,
		vecinos: []
	},
	{
		val: 65,
		vecinos: []
	},
	{
		val: 21,
		vecinos: []
	},
	{
		val: 18,
		vecinos: []
	}
]

//Se puede crear un array de un heap recorriendo el heap (arbol bonario) con el BFS:
function BFS(startNodeIndex, graph) {
	
	let queue = []
	queue.push(graph[startNodeIndex])
	let heapArray = []
	
	let visited = graph.map(() => false)
	visited[startNodeIndex] = true

	while (queue.length !== 0) {
		const nodoActual = queue.pop()
		heapArray.push(nodoActual.val)
		//console.log(nodoActual)
		const vecis = nodoActual.vecinos
		for (const vIndex of vecis) {
			if (!visited[vIndex]) {
				queue = [graph[vIndex], ...queue]
				visited[vIndex] = true
			}
		}
	}
	return heapArray;
}

console.log(BFS(0, heapTree))
