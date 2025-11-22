package main

import (
	"fmt"
	"os"
)

// processOrder simulates processing an order and uses defer for cleanup.
func processOrder(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // Ensure the file is closed even if an error occurs.

	// Process the order here.
	fmt.Println("Processing order from", filename)
	return nil
}

func main() {
	err := processOrder("order.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
