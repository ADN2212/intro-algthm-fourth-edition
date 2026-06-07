package main

import (
	"fmt"
	"math"
	"time"
)

// Este es el array de precios
// aqui p[i] = precio del rod de logitud i
var p [11]int = [11]int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}


var negInf int = int(math.Inf(-1))

// En el libro p pasa como argumento de la funcion, pero aqui dejo que se optenga por el scope,
// debido a que es el mismo para todos los casos
// esta funcion retorna la ganancia maxima que se le puede sacar a un rod de logitud n
//O(2^2) ver pagina 366
func CutRod(n int) int {
	if n == 0 {
		return 0
	}
	q := negInf
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

func MemoizedCutRod(n int) int {
	//Este es el array que se usa para memorizar las soluciones de los subtrees
	r := []int{}
	for i := 0; i <= n; i++ {
		r = append(r, negInf)
	}
	return memoizedCutRodAux(n, r)//como r es un slice pasa como una referencia.
}

//O(n^2)
func memoizedCutRodAux(n int, r []int) int {
	//Si la solucion esta memorizada se retorna, esto es lo que evita que se repitan calculos
	if r[n] >= 0 {
		return r[n]
	}
	var q int//Se inizializa como cero
	if n == 0 {
		return q
	} else {
		q = negInf
		for i:= 1; i <= n; i++ {
			q = max(q, p[i] + memoizedCutRodAux(n-i, r))
 		}
		r[n] = q//Esta es la parte en la que se memoriza la solucion
		return q
	}
}


func main() {
	fmt.Println("The Rod Cutting Problem")
	// fmt.Println("The max revenue for a rod of len 3 is = ", CutRod(3))
	// fmt.Println("The max revenue for a rod of len 4 is = ", CutRod(4))
	// fmt.Println("The max revenue for a rod of len 6 is = ", CutRod(6))

	// fmt.Println(MemoizedCutRod(4))

	//Aqui un ejemplo de la diferencia en performance estre la version normal la memo
	s1 := time.Now()
	fmt.Println(CutRod(10))
	e1 := time.Since(s1).Nanoseconds()
	//fmt.Println(e1)

	s2 := time.Now()
	fmt.Println(MemoizedCutRod(10))
	e2 := time.Since(s2).Nanoseconds()
	//fmt.Println(e2)

	fmt.Println("Memo was", e1 / e2, "times faster")

}
