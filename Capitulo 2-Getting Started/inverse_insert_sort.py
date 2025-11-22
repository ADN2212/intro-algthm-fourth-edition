#Correspondiente al ejercicio 2.1-3 en la pagina 23:

arr1 = [34, 12, 7, 56, 23, 89, 4, 71]
arr2 = [5, 29, 18, 2, 64, 31, 77, 10]
arr3 = [25, 90, 14, 6, 52, 83, 37, 19]


def reverse_insert_sort(arr):
	
	array = [n for n in arr]
	
	n = len(array)

	for i in range(1, n):
		
		key = array[i]
		j = i - 1
		print(f'Para key = {key}:')
		while j >= 0 and array[j] < key:#
			 print(f'		+ Moviendo {array[j]} a la posicion {j + 1}.')
			 array[j + 1] = array[j]			 
			 j = j - 1

		print(f'		- Moviendo {key} a la posicion {j + 1}.')
		array[j + 1] = key

	return array


#print(reverse_insert_sort(arr1))

print(arr2[0:0])



