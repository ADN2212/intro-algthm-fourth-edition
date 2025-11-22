function left(i) {
	return 2 * i
}

function rigth(i) {
	return 2 * i + 1
}

const heapArray2 = [null, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1]
const arr1 = [null, 8, 3, 5, 2, 10, 25]

//maxHeapyfy(heapArray2, 2)
//console.log(mainMaxHeapyfy(heapArray2, 2))
console.log(mainMaxHeapyfy(arr1, 1));

function mainMaxHeapyfy(heapArray, i) {
	const heapArrayCopy = heapArray.map(n => n)
	maxHeapyfy(heapArrayCopy, i)
	return heapArrayCopy
}

// Este algoritmo se encarga de hacer que un subtree de un heap optenga su max-heap property en caso de no tenerla. 
// ver pag 165.
function maxHeapyfy(heapArray, i) {
	const l = left(i)
	const r = rigth(i)
	let largest = -Infinity

	//En estas condicionales se asume que el array que se pasa solo tiene los elementos del heap.
	if (heapArray[l] > heapArray[i]) {
		largest = l
	} else {
		largest = i
	}

	if (heapArray[r] > heapArray[largest]) {
		largest = r
	}

	if (largest !== i) {
		console.log(`Cambiando ${heapArray[i]} por ${heapArray[largest]}.`)
		let temp = heapArray[i]
		heapArray[i] = heapArray[largest]
		heapArray[largest] = temp
		maxHeapyfy(heapArray, largest)
		return
	}

	return
}
