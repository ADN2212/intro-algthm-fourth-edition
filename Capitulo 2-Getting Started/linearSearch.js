//Ver ejercicio 2.1-4 del libro:
let arr2 = [5, 29, 18, 2, 64, 31, 77, 10];
let arr3 = [25, 90, 14, 6, 52, 83, 37, 19];

//Que Loops invariants pudes hallar en este array ?
//Para todo k tal que 1 < k <= n,

//se cumple si la funcion llega a la iteracion i = k,
//Entonces x no esta en el subarray A[1:k-1]
//Ojo: puede complicer que A[k] = x;

function lienarSearch(array, x) {
	for (let i = 0; i <= array.length; i++) {
		if (array[i] === x) return i
	}
	return null
}

console.log(lienarSearch(arr2, 5))
