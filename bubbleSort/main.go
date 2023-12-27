package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(arr []int, index int) {
	t := arr[index]
	arr[index] = arr[index+1]
	arr[index+1] = t
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
	return arr
}

func ConvertStringArrayToIntegerArr(arr []string) []int {
	var inputNumArr []int
	for i := 0; i < len(arr); i++ {
		numi, _ := strconv.Atoi(strings.TrimSpace(arr[i]))
		inputNumArr = append(inputNumArr, numi)
	}
	return inputNumArr
}

func main() {
	// array := []int{1, 3, 5, 4, 2}

	input := bufio.NewReader((os.Stdin))
	inputStr, _ := input.ReadString('\n')
	inputArr := strings.Split(inputStr, " ")

	inputArrNum := ConvertStringArrayToIntegerArr(inputArr)
	fmt.Println(BubbleSort(inputArrNum))
}
