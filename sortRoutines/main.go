package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ConvertStringArrayToIntegerArr(arr []string) []int {
	var inputNumArr []int
	for i := 0; i < len(arr); i++ {
		numi, _ := strconv.Atoi(strings.TrimSpace(arr[i]))
		inputNumArr = append(inputNumArr, numi)
	}
	return inputNumArr
}

func main() {
	var wg sync.WaitGroup

	// Prompt the user to input a series of integers
	fmt.Print("Enter integers (space-separated): ")

	input := bufio.NewReader((os.Stdin))
	inputStr, _ := input.ReadString('\n')
	inputArr := strings.Split(inputStr, " ")

	inputArrNum := ConvertStringArrayToIntegerArr(inputArr)

	fmt.Println(inputArrNum)
	// Calculate the partition sizes
	totalSize := len(inputArrNum)
	partitionSize := totalSize / 4

	fmt.Println(totalSize, ":", partitionSize)
	// Create channels for communication
	channels := make([]chan []int, 4)
	for i := range channels {
		channels[i] = make(chan []int)
	}

	// Print each subarray that will be sorted
	for i := 0; i < 4; i++ {
		startIndex := i * partitionSize
		endIndex := (i + 1) * partitionSize
		if i == 3 {
			endIndex = totalSize
		}

		subarray := inputArrNum[startIndex:endIndex]
		fmt.Printf("Goroutine %d will sort: %v\n", i+1, subarray)

		// 	// Launch goroutine to sort each subarray concurrently
		wg.Add(1)
		go func(i int, subarray []int) {
			defer wg.Done()
			sort.Ints(subarray)
			channels[i] <- subarray
		}(i, subarray)
	}

	// Close channels once all goroutines are done
	go func() {
		wg.Wait()
		for _, ch := range channels {
			close(ch)
		}
	}()

	// Merge the sorted subarrays
	mergedArray := mergeSortedArrays(<-channels[0], <-channels[1], <-channels[2], <-channels[3])

	// Print the entire sorted list
	fmt.Println("\nSorted Array:", mergedArray)
}

// mergeSortedArrays merges four sorted arrays into one

func mergeSortedArrays(arr1, arr2, arr3, arr4 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2)+len(arr3)+len(arr4))
	result = append(result, arr1...)
	result = append(result, arr2...)
	result = append(result, arr3...)
	result = append(result, arr4...)
	sort.Ints(result)
	return result
}
