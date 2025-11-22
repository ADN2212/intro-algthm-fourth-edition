import array

#Pair-1:
m1 = [
    [1, 1],
    [1, 1]
]

m2 = [
    [2, 2],
    [2, 2]
]

m3 = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12],
    [13, 14, 15, 16]
]

# Matriz B (4x4)
m4 = [
    [16, 15, 14, 13],
    [12, 11, 10, 9],
    [8, 7, 6, 5],
    [4, 3, 2, 1]
]

#Esta "corta" una matriz de la forma n x n con n % 2 = 0 en cuatro partes iaguales
def split_matrix(m):

    n = len(m)
    middle = (n // 2) - 1

    index = {
        "11": {
            "start_row": 0,
            "end_row": middle,
            "start_col":0,
            "end_col": middle,
        },
        "12": {
            "start_row": 0,
            "end_row": middle,
            "start_col": middle + 1,
            "end_col": middle + (n // 2)
        },
        "21": {
            "start_row": middle + 1,
            "end_row": middle + (n // 2),
            "start_col": 0,
            "end_col": middle
        },
        "22": {
            "start_row": middle + 1,
            "end_row": middle + (n // 2),
            "start_col": middle + 1,
            "end_col": middle + (n // 2)
        }
    }

    submatrices = {}

    for sm_str in ["11", "12", "21", "22"]: 
        curr_sm = []
        sr = index[sm_str]["start_row"]
        er = index[sm_str]["end_row"]

        for i in range(sr, er + 1):
            sc = index[sm_str]["start_col"]
            ec = index[sm_str]["end_col"]
            cur_row = m[i][sc:ec + 1]#Esto crea una nueva list, lo que proboca que las list dentro de submatrices no tengan refrencia a la matriz original. 
            curr_sm.append(cur_row)

        submatrices[sm_str] = curr_sm

    return submatrices

def show_matrix(m):
    for row in m:
        print(list(row))

#Asume que A y B tienen las mismas dimensiones.
def addMatrix(A, B, sing = 1):
    #Si sing = -1 estariamos sumando el opuesto, lo que es equivalente a restar.

    n = len(A)
    C = [[0] * n for _ in range(n)]

    #aqui se pueede ver claramente porque sumar dos matrices de n x n es O(n^2)
    for r in range(n):
        for c in range(n):
            C[r][c] = A[r][c] + (sing)* B[r][c]

    return C

#En todos lados asumo que todos las matrices son cuadradas y de la misma dimencion.
def merge(m1, m2, m3, m4):
    
    n = len(m1)
    res = []

    for i in range(n):
        res.append(m1[i] + m2[i])

    for j in range(n):
        res.append(m3[j] + m4[j])

    return res

def strassen_mian(A, B):
    #TODO: quitar strassen inner:
    def strassen_inner(a, b):
        #El caso base puede ser tambien de matrices 4x4 o dos por dos        
        #Caulquier caso cuyo resultado pueda ser descrito en un pequeño set de ecuaciones.
        if len(a) == 2:
            print("Hit the base case")
            c11 = a[0][0] * b[0][0] + a[0][1] * b[1][0]
            c12 = a[0][0] * b[0][1] + a[0][1] * b[1][1]
            c21 = a[1][0] * b[0][0] + a[1][1] * b[1][0]
            c22 = a[1][0] * b[0][1] + a[1][1] * b[1][1]

            return [
                [c11, c12],
                [c21, c22]
            ]

        print("Going to the recursive call")
        splited_a = split_matrix(a)
        splited_b = split_matrix(b)

        #Sumas de strassen, estas se ejecutan en O(n^2), O = theta
        s1 = addMatrix(splited_b["12"], splited_b["22"], sing = -1)
        s2 = addMatrix(splited_a["11"], splited_a["12"])
        s3 = addMatrix(splited_a["21"], splited_a["22"])
        s4 = addMatrix(splited_b["21"], splited_b["11"], sing = -1)
        s5 = addMatrix(splited_a["11"], splited_a["22"])
        s6 = addMatrix(splited_b["11"], splited_b["22"])
        s7 = addMatrix(splited_a["12"], splited_a["22"], sing = -1)
        s8 = addMatrix(splited_b["21"], splited_b["22"])
        s9 = addMatrix(splited_a["11"], splited_a["21"], sing = -1)
        s10 = addMatrix(splited_b["11"], splited_b["12"])
        
        #La P[i] de la pagina 87 corresponden a las llamadas recursivas:
        #Esta parte es la que hace que este algoritmo supere al de recursion normal
        #Ya que en vez de hacer 8 llamadas recursivas solo hace 7,
        #Lo que hace que a recurrence ecuation sea T(n) = 7 * T(n) + O(n^2)
        #Que al ser resuelta se optiene alguna O(n^lg(7)) lg = log base 2.
        p1 = strassen_inner(splited_a["11"], s1)#
        p2 = strassen_inner(s2, splited_b["22"])
        p3 = strassen_inner(s3, splited_b["11"])
        p4 = strassen_inner(splited_a["22"],  s4)
        p5 = strassen_inner(s5, s6)
        p6 = strassen_inner(s7, s8)
        p7 = strassen_inner(s9, s10)
        
        #Usar los Ps para calcular las sub-matrices:
        c11 = addMatrix(addMatrix(addMatrix(p5, p4), p2, sing = -1), p6)         
        c12 = addMatrix(p1, p2)
        c21 = addMatrix(p3, p4)
        c22 =  addMatrix(addMatrix(p5, p1), addMatrix(p3, p7), sing = -1)

        #mergear las matrices:
        return merge(c11, c12, c21, c22)

    return strassen_inner(A, B)


show_matrix(strassen_mian(m3, m4))
