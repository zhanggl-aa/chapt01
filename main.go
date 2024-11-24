package main

import (
	"fmt"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// http.ListenAndServe(":3000", nil)

	slice := []int{1, 22, 34}
	slice1 := slice
	for i, v := range slice1 {
		fmt.Println(v)
		if i == 1 {
			slice1[i] = 50
		}
	}
	fmt.Println(slice)

}
