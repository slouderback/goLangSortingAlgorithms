package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGraph(nums []int){
	for i := 0; i < len(nums); i++{
		for j := 0; j < nums[i]; j++{
			fmt.Print("#")
		}
		fmt.Print("\n")
		//fmt.Print(nums[i])
	}
	fmt.Print("\n\n")
}

func selectionSort(nums *[]int){
	for i := 0; i < len(*nums); i++{
		min := i
		for j := i+1; j < len(*nums); j++{
			if (*nums)[min] > (*nums)[j]{
				min = j
			}
		}
		if i == 0{
			fmt.Println("Initial condition:")
		} else{
			fmt.Printf("Step %d:\n", i)
		}
		printGraph(*nums)
		(*nums)[i], (*nums)[min] = (*nums)[min], (*nums)[i]
	}
}

func bubbleSort(nums *[]int){
	length := len(*nums)
	var swapped bool
	_ = swapped
	for i := 0; i < length-1; i++{
		if i == 0{
			fmt.Println("Initial condition:")
		} else{
			fmt.Printf("Step %d:\n", i)
		}
		printGraph(*nums)
		swapped = false
		for j := 0; j < length-i-1; j++{
			if (*nums)[j] > (*nums)[j+1]{
				(*nums)[j], (*nums)[j+1] = (*nums)[j+1], (*nums)[j]
				swapped = true
			}
		}
		if swapped == false{
			break;
		}
	}
}

func main(){
	nums := []int{7, 2, 10, 3, 8}
	fmt.Println("Order before sorting: ")
	printGraph(nums)


	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose sorting algorithm \n(1) Selection sort \n(2) Bubble sort\n")
	choice, _ := consoleReader.ReadString('\n')
	choice = strings.TrimRight(choice, "\n")
	if choice == "1" {
		selectionSort(&nums)
		fmt.Println("Order after selection sort: ")
		printGraph(nums)
	} else if choice == "2"{
		bubbleSort(&nums)
		fmt.Println("Order after bubble sort: ")
		printGraph(nums)
	}
}
