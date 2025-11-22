import array

#Esto hace que la matriz sea un array de punteros a los arrays, lo que me permitira usar las referencai a mi favor.
A = [
    memoryview(array.array('i', [1, 2, 3, 4])),
    memoryview(array.array('i', [5, 6, 7, 8])),
    memoryview(array.array('i', [9, 10, 11, 12])),
    memoryview(array.array('i', [13, 14, 15, 16]))
]

B = [
    memoryview(array.array('i', [7, 14, 3, 21])),
    memoryview(array.array('i', [5, 0, 11, 9])),
    memoryview(array.array('i', [8, 6, 13, 4])),
    memoryview(array.array('i', [2, 19, 1, 17]))
]

#Esta "corta" una matriz de la forma n x n con n % 2 = 0 en cuatro partes iaguales
def split_matrix(m):

    n = len(m)
    middle = (n // 2) - 1

    index = {
        "A11": {
            "start_row": 0,
            "end_row": middle,
            "start_col":0,
            "end_col": middle,
        },
        "A12": {
            "start_row": 0,
            "end_row": middle,
            "start_col": middle + 1,
            "end_col": middle + (n // 2)
        },
        "A21": {
            "start_row": middle + 1,
            "end_row": middle + (n // 2),
            "start_col": 0,
            "end_col": middle
        },
        "A22": {
            "start_row": middle + 1,
            "end_row": middle + (n // 2),
            "start_col": middle + 1,
            "end_col": middle + (n // 2)
        }
    }

    submatrices = {}

    for sm_str in ["A11", "A12", "A21", "A22"]: 
        curr_sm = []
        sr = index[sm_str]["start_row"]
        er = index[sm_str]["end_row"]

        for i in range(sr, er + 1):
            sc = index[sm_str]["start_col"]
            ec = index[sm_str]["end_col"]
            cur_row = m[i][sc:ec + 1]
            curr_sm.append(cur_row)

        submatrices[sm_str] = curr_sm

    return submatrices


#Esta funcion asume que la matriz entante es cuadrada de con n multiplo de dos:
#theta(n^3)
def recursive_multiply_matrix(a, b):

    n = len(a)
    res = [memoryview(array.array('i', [0] * n)) for i in range(n)]

    def inner_recursive_multiply(m1, m2, c):
        
        if len(m1) == 1:
            c[0][0] = c[0][0] + m1[0][0] * m2[0][0]
            return
        
        splited_m1 = split_matrix(m1)
        splited_m2 = split_matrix(m2)
        splited_c = split_matrix(c)


        inner_recursive_multiply(m1 = splited_m1["A11"], m2 = splited_m2["A11"], c = splited_c["A11"])
        inner_recursive_multiply(m1 = splited_m1["A11"], m2 = splited_m2["A12"], c = splited_c["A12"])

        inner_recursive_multiply(m1 = splited_m1["A21"], m2 = splited_m2["A11"], c = splited_c["A21"])
        inner_recursive_multiply(m1 = splited_m1["A21"], m2 = splited_m2["A12"], c = splited_c["A22"])
        
        inner_recursive_multiply(m1 = splited_m1["A12"], m2 = splited_m2["A21"], c = splited_c["A11"])
        inner_recursive_multiply(m1 = splited_m1["A12"], m2 = splited_m2["A22"], c = splited_c["A12"])

        inner_recursive_multiply(m1 = splited_m1["A22"], m2 = splited_m2["A21"], c = splited_c["A21"])        
        inner_recursive_multiply(m1 = splited_m1["A22"], m2 = splited_m2["A22"], c = splited_c["A22"])

    inner_recursive_multiply(m1 = a, m2 = b, c = res)
    return res


m = recursive_multiply_matrix(A, B)

for row in m:
    print(list(row))
