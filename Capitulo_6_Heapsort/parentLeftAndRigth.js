//El null al principio se utiliza para llenar la primera posicion y hacer que el array "empiece" desde cero.
const heapArray = [null, 90, 89, 70, 36, 75, 63, 65, 21, 18]
const heapArray2 = [null, 33, 19, 20, 15, 13, 10, 2, 13, 16, 12]

function parent(i) {
 return Math.floor(i / 2)
}

function left(i) {
 return 2*i
}

function rigth(i) {
 return 2*i + 1
}

//console.log(heapArray[parent(4)])
//console.log(heapArray[left(2)])
//console.log(heapArray[rigth(3)])

//Vease como los elementos del heap cumplen con la max-heap property:
//A[parent(i)] >= A[i], para i >= 2, porque A[1] no tiene padre:

for (let i = 2; i < heapArray2.length; i++) {
 //console.log(i)
 console.log(`A[${parent(i)}] >= A[${i}] = ${heapArray2[parent(i)] >= heapArray2[i]}`)
}
