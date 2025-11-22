const heapSort = (function () {

    function left(i) {
        return 2 * i
    }

    function rigth(i) {
        return 2 * i + 1
    }

    function buildMaxHeap(arr) {

        const maxHeap = arr.map(n => n);
        const maxHeapLen = maxHeap.length - 1;
        let startIndex = Math.floor((maxHeap.length - 1) / 2)
        //console.log(startIndex);
        
        for (startIndex; startIndex > 0; startIndex--) {
            //console.log(startIndex)
            maxHeapyfy(maxHeap, startIndex, maxHeapLen);
        }

        return maxHeap;
    }

    function maxHeapyfy(heapArray, i, heapLen) {
        const l = left(i)
        const r = rigth(i)
        let largest = -Infinity

        //Esto es para imitar la linea 4 del algoritmo.
        if (i >= heapLen) {
            console.log(`i = ${i}, heapLen = ${heapLen}.`)
            return
        }

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
            //console.log(`En i = ${i}, Cambiando ${heapArray[i]} por ${heapArray[largest]}.`)
            let temp = heapArray[i]
            heapArray[i] = heapArray[largest]
            heapArray[largest] = temp
            maxHeapyfy(heapArray, largest, heapLen)
            return
        }

        return
    }

    function heapSortMain(array) {
        const maxHeap = buildMaxHeap(array);
        //console.log(maxHeap);
        const n = maxHeap.length - 1;
        let heapLen = n;
        //console.log(n)
        let tail = null;
        let head = null;
        //return false
        for (let i = n; i >= 2; i--) {
            //console.log(i)        
            head = maxHeap[1]
            tail = maxHeap[i];
            //console.log(`Cambiando A[1] = ${head} con A[${i}] = ${tail}`) 
            maxHeap[1] = tail
            maxHeap[i] = head
            heapLen = heapLen - 1
            maxHeapyfy(maxHeap, 1, heapLen)
            //console.log(maxHeap);
        }
        return maxHeap;
    }

    return heapSortMain;

})()

const arr1 = [null, 8, 3, 5, 2, 10, 25]

console.log(heapSort(arr1));
