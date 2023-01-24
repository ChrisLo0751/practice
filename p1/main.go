package main

import (
	"fmt"
	"os"
	 "github.com/practice/p1/pkg"
)


func main() {
	var size int
	var key string
	var val int

	fmt.Println("Please enter you hope queue size")
	fmt.Scanln(&size)
	fmt.Println("Wait...initial queue")
	var q = pkg.NewSingleQueue(size)
	fmt.Println("Successfully initial queue")
	
	for {
		fmt.Println("1. Enter add to add data to the queue")
		fmt.Println("2. Enter get to get the queue")
		fmt.Println("3. Enter show to show the queue")
		fmt.Println("4. Enter exit to exit")

		fmt.Scanln(&key)
		switch key {
			case "add":
			fmt.Println("Enter the data you want to join the queue")
			fmt.Scanln(&val)
			err := q.AddQueue(val)
			if err != nil {
				fmt.Printf("failed to add queue, err: %v\n", err)
			}else{
				fmt.Printf("Successfully added %d \n", val)
				q.ShowQueue()
			}
			break
			case "get":
			case "show":
				q.ShowQueue()
			case "exit":
				os.Exit(0)
			default:
			fmt.Println("Please enter correct value")
		}
	}
}