function left(i) {
	return 2 * i
}

function rigth(i) {
	return 2 * i + 1
}

//const heapArray2 = [null, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1]
//maxHeapyfy(heapArray2, 2)
//console.log(mainMaxHeapyfy(heapArray2, 2))

const array = [null, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7]
const array2 = [null, 5, 3, 17, 10, 84, 19, 6, 22, 9]

//Esta funcion construye un maxHeap con el array que toma como input, ver pagina 167
function buildMaxHeap(arr) {

    const maxHeap = arr.map(n => n);
    let startIndex = Math.floor((maxHeap.length - 1) / 2)
	console.log(startIndex);

    for (startIndex; startIndex > 0; startIndex--) {
        //console.log(startIndex)
        maxHeapyfy(maxHeap, startIndex);
    }

    return maxHeap;

}

console.log(buildMaxHeap(array2));

// function mainMaxHeapyfy(heapArray, i) {
// 	const heapArrayCopy = heapArray.map(n => n)
// 	maxHeapyfy(heapArrayCopy, i)
// 	return heapArrayCopy
// }

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
		console.log(`En i = ${i}, Cambiando ${heapArray[i]} por ${heapArray[largest]}.`)
		let temp = heapArray[i]
		heapArray[i] = heapArray[largest]
		heapArray[largest] = temp
		maxHeapyfy(heapArray, largest)
		return
	}

	return
}
