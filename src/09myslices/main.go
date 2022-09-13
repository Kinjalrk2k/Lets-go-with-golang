package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcom to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	// fmt.Println("Type of fruitList is %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// used for slicing
	fruitList = (fruitList[1:3]) // [1:] or [:3] also works
	fmt.Println(fruitList)

	// define a slice with make
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// highScores[4] = 777 // crashed
	highScores = append(highScores, 555, 666, 777) // doesn't crash - append reallocated memory
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // works in place
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println("----------------------------------")

	// remove a value from slice
	var courses = []string{"react", "javascript", "swift", "python", "go"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
