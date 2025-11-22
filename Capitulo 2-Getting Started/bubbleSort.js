let arr1 = [5, 2, 4, 6, 1, 3];
let arr2 = [91, 2, 76, 18, 54, 33];
let arr3 = [3, 41, 52, 26, 38, 57, 9, 49];

function bubbleSort(array) {

	let sortedArray = array.map(n => n)
	const n = array.length

	for (let i = 0; i < n; i++) {
		//console.log(`Para i = ${i}:`)
		for (let j = n - 1; j >= i + 1; j--) {
			console.log(`j = ${j}, i = ${i}`)
			//console.log(`	array[j = ${j}] = ${sortedArray[j]}.`)
			//console.log(`	array[j - 1 = ${j - 1}] = ${sortedArray[j - 1]}.`)
			if (sortedArray[j] < sortedArray[j - 1]) {
				exchange(sortedArray, j, j - 1)
			}
			//console.log(`Al final de ciclo secundario => ${sortedArray}`)
		}
		//console.log(`Al final de ciclo principal => ${sortedArray}`)
	}
	return sortedArray
}

//Esta funcion cambia pone array[i] donde esta array[j] y viseversa:
//Toma el array como referencai, por lo que funciona como un metodo.
function exchange(array, i, j) {
	//console.log(`	Cambiando array[${i}] = ${array[i]} con array[${j}] = ${array[j]}.`)
	const aux = array[i]
	array[i] = array[j]
	array[j] = aux
	return
}

console.log(bubbleSort(arr2))
