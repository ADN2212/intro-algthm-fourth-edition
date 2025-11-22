#Estos son los coeficientes del polinomio a0 + a1x, a2(x^2) + ... + an(x^n)

coefficients_1 = [5, 3, 2]
coefficients_2 = [10, 2, 4, -7]

#Este procedimiento imita la expresion: a0 + x(a1 + x(a2 + ... + x(an-1 + x an)))
#para calcular el valor de un polinomio P(x), ver problema 2-3 en la pagina 46-47.

def honer(cs, x):
	p = 0
	n = len(cs)
	for i in range(n):
		print(f'p = {p} at the starting of i = {i}')
		j = n - (i + 1)
		p = cs[j] + (x * p)
	return p

print(honer(coefficients_1, 3))
#print(honer(coefficients_2, 2))


def polynomial_evaluation(cs, x):
	res = 0
	for i in range(len(cs)):
		#print(f'{cs[i]} por {x}^{i}')
		res += cs[i] * pow(x, i)
	return res

#print(polynomial_evaluation(coefficients_1, 2))
#print(polynomial_evaluation(coefficients_2, 2))
