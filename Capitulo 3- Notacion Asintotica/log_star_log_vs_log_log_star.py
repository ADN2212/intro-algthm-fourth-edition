from math import log2

lg = lambda n : log2(n)

#Ver la definicion de lg* n en la pagina 68.
#Esta funcion cuenta la veces que hay que aplicar lg para hacer
#que el resultado sea menor o igual que 1.  
def log_star(n):
	
	if lg(n) <= 1:
		return 1

	return 1 + log_star(lg(n))

#print(log_star(16))


def log_log_star(n):
	return lg(log_star(n))

def log_star_log(n):
	return log_star(lg(n))


for i in range(2, 1000):
	
	#print(f'For {i} is {log_star(i)}')

	lg1 = log_log_star(i)
	lg2 = log_star_log(i)

	print(f'{i},{lg1},{lg2}')

	# if lg1 > lg2:
	# 	print(f'log log star is greater for {n}')
	# elif lg2 > lg1:
	# 	print(f'log star log is greater for {n}')
	# else :
	# 	print(f'They are equal for {n}')

