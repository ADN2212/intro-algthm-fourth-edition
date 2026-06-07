package main

import "fmt"
//import "math"

// Este es el array de precios
// aqui p[i] = precio del rod de logitud i
var p [11]int = [11]int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

// En el libro p pasa como argumento de la funcion, pero aqui dejo que se optenga por el scope,
// debido a que es el mismo para todos los casos
// esta funcion retorna la ganancia maxima que se le puede sacar a un rod de logitud n
//O(2^2) ver pagina 366
func CutRod(n int) int {
	if n == 0 {
		return 0
	}
	q := -1000000//int(math.Inf(-1))
	for i := 1; i <= n; i++ {
		//En esta vercion la recursion hace que se repita llamadas, ver la pagina 367
		q = max(q, p[i] + CutRod(n - i))
	}
	return q
}

//Esto es solo para que se vea como en el pseudocodigo del libro
func max(a, b int) int {
	if a > b {
		return a
	} 
	return b
}


func main() {
	fmt.Println("The Rod Cutting Problem")
	fmt.Println("The max revenue for a rod of len 3 is = ", CutRod(3))
	fmt.Println("The max revenue for a rod of len 4 is = ", CutRod(4))
	fmt.Println("The max revenue for a rod of len 6 is = ", CutRod(6))

}
