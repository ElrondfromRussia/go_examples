package main

import (
	"fmt"
	"strconv"
)

func main(){
	var ss string
	fmt.Scan(&ss)
	var n int
	fmt.Scan(&n)
	var i int = 0;
	for i < len(string(ss)){
		if string(ss[i]) == strconv.Itoa(n){
			ss = string(ss[:i]) +  string(ss[i+1:])
			continue
		}
		i++
	}
	fmt.Println(ss)
}
