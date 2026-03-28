
package main

import "fmt"

func main() {
	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	sate5 := 0
	smais5 := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] <= 5 {
			sate5 += slice[i]
		} else {
			smais5 += slice[i]
		}
	}

	fmt.Println(sate5)
	fmt.Println(smais5)

}
