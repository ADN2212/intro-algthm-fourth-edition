#Implementacion del algoritmo de multiplicacion de matrices (ver pag 81)
m1 = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]

# Matriz 3x3 
m2 = [
    [9, 8, 7],
    [6, 5, 4],
    [3, 2, 1]
]

m3 = [
    [2, 3],
    [6, 7]
]

m4 = [
    [10, 2],
    [5, 4]
]

#Theta(n^3)
def matrix_multiply(A, B):
    #Asumo que A y B son cuadradas y me ahorro hacer validacines ya que este codigo es "academico"
    n = len(A)

    #inizializar C:
    C = []#No hagas [[0] * n] * n pa que los valores por referencia no te jodan el dia.

    for a in range(n):
        C.append([0] * n)

    for i in range(n):
        for j in range(n):
            for k in range(n):
                print(f'C[{i}][{j}] = {C[i][j]} + {A[i][k]} * {B[k][j]}')
                C[i][j] = C[i][j] + A[i][k] * B[k][j]#Ver ecuacion 4.1 en la pagina 81
                
    return C

def show_matrix(m):
    for row in m:
        print(row)

show_matrix(matrix_multiply(m3, m4))
