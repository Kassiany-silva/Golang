\*Crie um array de duas posições e some seus valores*\


package main

import "fmt"

func main() {
	arr := []int{4, 5}
	soma := 0
	for i := 0; i < len(arr); i++ {
		soma += arr[i]
	}

	fmt.Println(soma)
}
