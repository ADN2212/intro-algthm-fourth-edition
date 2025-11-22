let arr1 = [5, 2, 4, 6, 1, 3];
let arr2 = [91, 2, 76, 18, 54, 33];
let arr3 = [3, 41, 52, 26, 38, 57, 9, 49];
let arr4 = [1];

//Ver ejercicio
const inSort = (function () {
	
	function inSortMain(arr) {
		if (arr.length === 0) return []; //Este es el caso espesial de un array sin elementos
		let array = arr.map((n) => n);
		innerInSort(1, 0, array[1], array);
		return array;
	}

	function innerInSort(i, j, key, array) {
		if (i === array.length) {
			return;
		}

		if (j >= 0 && array[j] > key) {
			array[j + 1] = array[j];
			innerInSort(i, j - 1, key, array);
			return;
		}

		array[j + 1] = key;
		innerInSort(i + 1, i, array[i + 1], array);
		return;
	}

	return inSortMain;

})();

//console.log(inSortMain(arr1))
//console.log(inSortMain(arr2))
console.log(inSort(arr4));
