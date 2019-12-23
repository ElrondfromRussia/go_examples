package main

import (
	"fmt"
)


func fib(n int) int{
	var a = 0
	var b = 1
	var k = 0
	for k < n {
		a, b = b, a+b;
		k++;
	}
	return a
}


func main(){
	var n int
	fmt.Scan(&n)
	var sum int = -1;
	var fb int = 0;
	var i int = 1;
	for fb < n{
		fb = fib(i);
		if fb == n{
			sum = i;
			break;
		}
		i++;
	}
	fmt.Println(sum)
}
