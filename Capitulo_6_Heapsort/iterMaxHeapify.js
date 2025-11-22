function left(i) {
    return 2 * i
}

function rigth(i) {
    return 2 * i + 1
}

const heapArray2 = [null, 16, 4, 10, 14, 7, 9, 3, 2, 8, 1]

console.log(iterMaxHeapify(heapArray2, 2));

function iterMaxHeapify(heapArray, i) {

    const heapArrayCopy = heapArray.map(n => n);
    let currentLeft = null;
    let currentRigth = null;
    let currentlargest = 0;

    while (true) {

        currentLeft = left(i)
        currentRigth = rigth(i)

        if (heapArrayCopy[currentLeft] > heapArrayCopy[i]) {
            currentlargest = currentLeft
        } else {
            currentlargest = i
        }

        if (heapArrayCopy[currentRigth] > heapArrayCopy[currentlargest]) {
            currentlargest = currentRigth
        }

        if (currentlargest !== i) {
            console.log(`Cambiando ${heapArrayCopy[i]} por ${heapArrayCopy[currentlargest]}.`)
            let temp = heapArrayCopy[i]
            heapArrayCopy[i] = heapArrayCopy[currentlargest]
            heapArrayCopy[currentlargest] = temp
            i = currentlargest
        } else {
            break
        }
    }

    return heapArrayCopy

}



