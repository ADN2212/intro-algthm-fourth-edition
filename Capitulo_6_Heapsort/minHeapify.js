const heapArray3 = [null, 1, 10, 3, 5, 2]
const heapArray4 = [null, 1, 4, 20, 15, 10, 5]

//console.log(minHeapify);

var minHeapify = (function () {

  function left(i) {
    return 2 * i
  }

  function rigth(i) {
    return 2 * i + 1
  }

  function innerMinHeapify(heapArray, i) {

    const l = left(i)
    const r = rigth(i)
    let smallest = null;

    if (heapArray[l] < heapArray[i]) {
      smallest = l
    } else {
      smallest = i
    }

    if (heapArray[r] < heapArray[smallest]) {
      smallest = r
    }

    if (smallest !== i) {
      let temp = heapArray[i]
      heapArray[i] = heapArray[smallest]
      heapArray[smallest] = temp
      innerMinHeapify(heapArray, i)
      return
    }
  }

  return function main(heapArray, i) {
    const heapArrayCopy = heapArray.map(n => n)
    innerMinHeapify(heapArrayCopy, i)
    return heapArrayCopy
  }
    ;
})()

//console.log(minHeapify)
//console.log(minHeapify(heapArray3, 2))
console.log(minHeapify(heapArray4, 3))
