a1 = [34, 12, 7, 89, 45, 23]

a2 = [91, 2, 76, 18, 54, 33]

a3 = [47, 85, 29, 14, 66, 3]

a4 = [63, 27, 81, 5, 92, 38]

a5 = [9, 74, 41, 57, 16, 88]

bests_case = [1, 2, 3, 4, 5, 6, 7]

worts_case = [6, 5, 4, 3, 2, 1, 8]

#print(a1[0:0])

def selection_sort(a):
	array = [n for n in a]
	current_min = float("inf")
	current_min_pos = None
	array_len = len(array)
	#outter_loop_counter = 0
	#inner_loop_counter = 0

	for i in range(array_len - 1):#ver las notas para saber porque solo hay que hacer n - 1 iteraciones.
		#print(f'Ín iteration number: {i}')
		#outter_loop_counter += 1
		current_sub_array = array[i:array_len]
		#print(current_sub_array)
		for j in range(len(current_sub_array)):
			#inner_loop_counter += 1
			if current_sub_array[j] < current_min:
				current_min =  current_sub_array[j]
				current_min_pos = i + j
				#print(f'New Min: {current_min} at pos: {current_min_pos}')
		#print(f'Intercambiando {current_min} y {array[i]}.')
		#print(inner_loop_counter)
		inner_loop_counter = 0
		array[current_min_pos] =  array[i]
		array[i] = current_min
		current_min = float("inf")
		#print(f'Array = {array} after i = {i + 1}.')

	# print(outter_loop_counter)
	# print(inner_loop_counter)
	#print(f'Number of iterations = {outter_loop_counter * inner_loop_counter}')
	return array

print(selection_sort(a1))
#print(selection_sort(a2))
#print(selection_sort(a3))
#print(selection_sort(a4))
#print(selection_sort(bests_case))
#print(selection_sort(worts_case))