package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func printAnswers(answers []int, index int) int {
	if index == len(answers){
		return 0
	}
	fmt.Println(answers[index])
	return printAnswers(answers, index + 1)
}

func summer(nums []string, index int) int {
	if index < 0 {
		return 0
	}

	num, _ := strconv.Atoi(nums[index])

	if num < 0 {
		return summer(nums, index - 1)
	} else {
		
		return num * num + summer(nums, index - 1)
	}
	
}

func cases(numOfCases int, answers []int) []int {
	if numOfCases == 0 {
		return answers
	}
	
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	numOfNums, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Invalid Input")
		os.Exit(1)
	}

	text2, _ := reader.ReadString('\n')
	text2 = strings.Replace(text2, "\r\n", "", -1)
	nums := strings.Split(text2, " ")
	
	
	if numOfNums != len(nums) {
		fmt.Println("Invalid Input")
		os.Exit(1)
	}

	answers = append(answers, summer(nums, numOfNums - 1))
	
	return cases(numOfCases - 1, answers)
	
}


func main () {

	reader := bufio.NewReader(os.Stdin)
	
	numOfCases := 0

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r\n", "", -1)
	numOfCases, _ = strconv.Atoi(text)
	var answers []int

	answers = cases(numOfCases, answers)
	
	index := 0

	printAnswers(answers, index)	

	os.Exit(1)
		
	
}