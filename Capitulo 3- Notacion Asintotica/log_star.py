from math import log2

lg = lambda n : log2(n)

#Ver la definicion de lg* n en la pagina 68.
#Esta funcion cuenta la veces que hay que aplicar lg para hacer
#que el resultado sea menor o igual que 1.  
def log_star(n):
	
	if lg(n) <= 1:
		return 1

	return 1 + log_star(lg(n))

print(lg(65536))
print(log_star(3))
