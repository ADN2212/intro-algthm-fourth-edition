let arr1 = [5,2,4,6,1,3]

function insertionSort(a) {

	let n = a.length
	let array = a.map(n => n)//Para mantener el algoritmo como una funcion pura.

	for (let i = 1; i < n; i++) {
		let key = array[i]
		console.log(`key = ${key}`)
		let j = i - 1
		while (j >= 0 && array[j] > key) {
			console.log(`+Moviendo ${array[j]} a la posicion ${j + 1}`)
			array[j + 1] = array[j]
			j = j - 1
		}
		console.log(`-Moviendo ${key} a la posicion ${j + 1}`)	
		array[j + 1] = key
		console.log(`Array actual = ${array}`)
	}
	return array
}

console.log(insertionSort(arr1))
//console.log(arr1)
