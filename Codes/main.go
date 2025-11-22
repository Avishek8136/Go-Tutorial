package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	println("Create a new bill")
	print("Enter bill ID: ")
	// idinput:=bufio.NewReader(os.Stdin)
	// idinput,_ = idinput.ReadString('\n')
	// idinput = strings.TrimSpace(idinput)
	idInput := getInput("Enter bill ID: ", reader)
	id, _ := strconv.Atoi(idInput)
	bill := NewBill(id)
	return bill
}

func promptOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt := getInput("Choose option (a - add item, s -save bill, t - add tip) ", reader)
	switch opt {
	case "a":
		itemName := getInput("Enter item name: ", reader)
		itemPriceInput := getInput("Enter item price: ", reader)
		itemPrice, err := strconv.ParseFloat(itemPriceInput, 64)
		if err != nil {
			fmt.Println("Invalid price. Please try again.")
			promptOptions(b)
			return
		}
		b.addItem(itemName, itemPrice)
		fmt.Println("Item added successfully!")
		promptOptions(b)
	case "t":
		tipInput := getInput("Enter tip amount: ", reader)
		tip, err := strconv.ParseFloat(tipInput, 64)
		if err != nil {
			fmt.Println("Invalid tip amount. Please try again.")
			promptOptions(b)
			return
		}
		b.updateTip(tip)
		fmt.Println("Tip updated successfully!")
		promptOptions(b)
	case "s":
		fmt.Println("Saving bill...")
		fmt.Println(b.format())
	default:
		fmt.Println("Invalid option. Please try again.")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}

//{
// myBill.updateTip(2.5)
// myBill.addItem("coffee", 3.99)
// println(myBill.format())
// fmt.Println(myBill)
// fmt.Printf("Type: %T\n", myBill)
// fmt.Println("Bill ID:", myBill.ID)
// fmt.Println("Bill Items:", myBill.Items)
//}

// func updateName(x *string) {
// 	*x = "hammer"
// }

// func main() {
// 	name := "tool"

// 	updateName(&name)
// 	fmt.Println(name)
// }

// func updateName(x string) string {
// 	x = "hammer"
// 	return x
// }

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

// func main() {
// 	// group A types -> strings, ints, bools, floats, arrays, structs
// 	name := "tool"

// 	name = updateName(name)
// 	fmt.Println(name)

// 	// group B types -> slices, maps, functions
// 	menu := map[string]float64{
// 		"pi":     3.41,
// 		"banana": 8,
// 	}

// 	updateMenu(menu)
// 	fmt.Println(menu)
// }

// // maps
// func main() {
// 	person := map[string]string{
// 		"name":    "Alice",
// 		"country": "Wonderland",
// 		"job":     "Adventurer",
// 	}
// 	fmt.Println("Person map:", person["name"])
// 	fmt.Printf("Type: %T\n", person)
// 	fmt.Printf("Length: %d\n", len(person))
// 	fmt.Println(person)
// 	for key, value := range person {
// 		fmt.Printf("%s: %s\n", key, value)
// 	}
// }

//package scope main

// func main() {
// 	sayHello("World!")
// 	println(a)
// }

// func tryMe(message string) (string, string) {
// 	return message, "GoLang"
// }
// func main() {
// 	a, b := tryMe("Hello, World!")
// 	print(a, b)
// }

// func myfunc() {
// 	fmt.Println("This is my function")
// }

// func myfunc(message string) {
// 	fmt.Println(message)
// 	return 0,4
// }

// func tryMe(message string) {
// 	fmt.Println(message)
// 	myfunc(message)
// }

// func tryMeAgain(message string, f func(string)) {
// 	f(message)
// }

// func main() {
// 	tryMe("Hello from main")
// 	tryMeAgain("Hello from tryMeAgain", myfunc)
// 	tryMeAgain("Greetings", tryMe)
// }

// func main() {
// 	names := []string{"Alice", "Bob", "Charlie"}
// 	for index, name := range names {
// 		if index == 1 {
// 			continue
// 		} else {
// 			fmt.Printf("Index: %d, Name: %s\n", index, name)
// 		}
// 	}

//booleans
// age := 20
// if age >= 18 {
// 	fmt.Println("You are an adult.")
// }else if age == 18 {
// 	fmt.Println("You are just turned adult.")
// } else {
// 	fmt.Println("You are a minor.")
// fmt.Println(age >= 18)
// fmt.Println(age <= 18)
// fmt.Println(age == 18)
// fmt.Println(age != 18)
//}

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
