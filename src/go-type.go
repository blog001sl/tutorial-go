package main

import "fmt"

func main2() {
	var numbers2 [5]int
	numbers2[0] = 1
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)

	sum := (9)

	numbers1 := [...]int{1, 0, -2, 5, 5}

	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
	fmt.Printf("%v\n", (sum == numbers1[0]+numbers1[1]+numbers1[2]+numbers1[3]+numbers1[4]))

	slice1 := numbers1[1:4]
	capS1 := cap(slice1)
	var slice3 []int
	capS2 := cap(slice3)

	fmt.Printf("cap s1 is %d length s1 is %v, cap s2 is %d, length s2 is %v", capS1, len(slice1), capS2, len(slice3))

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice5 := numbers4[4:6:8]
	fmt.Printf("\n%v, %v, %v\n", slice5, len(slice5), cap(slice5))
	slice5 = slice5[:cap(slice5)]
	fmt.Printf("\n%v, %v, %v\n", slice5, len(slice5), cap(slice5))
	slice5 = append(slice5, 11, 12, 13, 14, 15)
	fmt.Printf("\n%v, %v, %v\n", slice5, len(slice5), cap(slice5))
	slice6 := []int{20, 21, 22}
	copy(slice5, slice6)
	fmt.Printf("\n%v, %v, %v\n", slice5, len(slice5), cap(slice5))
	fmt.Printf("\n%v\n", numbers4)

	map1 := map[string]int{"one": 1, "two": 2}
	fmt.Printf("%v\n", map1)

	mv1, ok := map1["one"]
	fmt.Printf("%v[%v]\n", mv1, ok)

	mv1, ok = map1["tow"]
	fmt.Printf("%v[%v]\n", mv1, ok)

	ch2 := make(chan string, 5)
	go func() {
		ch2 <- "channel value 1"
	}()

	ch2value := <-ch2
	fmt.Printf("%v\n", ch2value)
}
