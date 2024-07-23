package arrayslicemap

import "fmt"

func SliceMake() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)
	fmt.Println("---------------------------------------------------------\n")

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
	fmt.Println("---------------------------------------------------------\n")

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
	fmt.Println("---------------------------------------------------------\n")

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
