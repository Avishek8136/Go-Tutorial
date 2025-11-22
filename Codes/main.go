package main

func main() {
}

//loops
// x := 0
// for x < 5 {
// 	fmt.Println("Value of x:", x)
// 	x++
// }

// for i := 0; i < 5; i++ {
// 	fmt.Println("Value of i:", i)
// }

// names := []string{"Alice", "Bob", "Charlie"}
// for i := 0; i < len(names); i++ {
// 	fmt.Printf("Index: %d, Name: %s\n", i, names[i])
// }

// names := []string{"Alice", "Bob", "Charlie"}
// for i, name := range names {
// 	fmt.Printf("Index: %d, Name: %s\n", i, name)
// }

//}

// func main() {
// 	names := []string{"Alice", "Bob", "Charlie"}
// 	sort.Strings(names)
// 	fmt.Println("Sorted names:", names)
// 	fmt.Println("Index of 'Bob':", sort.SearchStrings(names, "Bob"))
// 	sort.Sort(sort.Reverse(sort.StringSlice(names)))
// 	fmt.Println("Reverse sorted names:", names)
// }

// func main() {
// 	//sort package
// 	numbers := []int{5, 2, 6, 3, 1, 4}
// 	fmt.Println("Before sorting:", numbers)
// 	sort.Ints(numbers)
// 	fmt.Println("After sorting:", numbers)

// 	index := sort.SearchInts(numbers, 10)
// 	fmt.Println("Index of 3 is:", index)
// }

// func main() {
// 	//strings
// 	greeting := "Hello, World!"
// 	fmt.Println(strings.Contains(greeting, "World"))
// 	fmt.Println(strings.ReplaceAll(greeting, "World", "Go"))
// 	fmt.Println(strings.ToUpper(greeting))
// 	fmt.Println(strings.Index(greeting, "X"))
// 	fmt.Println(strings.Split(greeting, ", "))
// 	fmt.Println(strings.HasPrefix(greeting, "Hello"))
// }

// func main() {
// 	//slices
// 	var scores = []int{90, 80, 70, 60, 50}
// 	scores[1] = 85
// 	scores = append(scores, 40, 30)
// 	fmt.Println(scores)
// 	fmt.Printf("Type: %T\n", scores)
// 	fmt.Print(len(scores))
// 	val := scores[2:5]
// 	fmt.Println("\n", val)
// }

// func main() {
// 	//Arrays and slices
// 	//var arr [5]int = [5]int{1, 2, 3, 4, 5}
// 	//var ages = [3]int{10, 20, 30}
// 	ages := [5]int{10, 20, 30, 40, 50}
// 	fmt.Println(ages)
// 	fmt.Printf("Type: %T\n", ages)
// 	fmt.Print(len(ages))
// }

// func main() {
// 	var salary = 25.00
// 	fmt.Printf("Salary is %f\n", salary)
// 	var st = fmt.Sprintf("Salary is %f\n", salary)
// 	fmt.Println(st)
// }

// func main() {
// 	//strings
// 	var name string = "World"
// 	var name1 = "Hello"
// 	fmt.Print(name1 + ", ")
// 	fmt.Print(name)
// 	fmt.Println("!")
// 	fmt.Printf("%T,", name)
// 	fmt.Printf("%T", name1)
// }

// 	var name3 string
// 	name2 := "!"
// 	var valu int8 = 25
// 	fmt.Println(name1 + ", " + name + name2 + name3)
// 	fmt.Println("The value is", valu)
// 	fmt.Printf("The value is %v\n", valu)
// 	fmt.Println(name1, name2, name3)
// }
