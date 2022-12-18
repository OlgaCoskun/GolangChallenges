package main

import "fmt"

func main() {
	var bilet string
	fmt.Scan(&bilet)

	if bilet[0] + bilet[1] + bilet[2] == bilet[3] + bilet[4] + bilet[5] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере 
которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных

На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
*/