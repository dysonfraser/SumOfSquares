package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// Function to print out the summed numbers from all the cases
func printAnswers(answers []int, index int) int {
	if index == len(answers){
		return 0
	}
	fmt.Println(answers[index])
	return printAnswers(answers, index + 1)
}


// Function to square postive numbers and sum them while ignoring negative numbers
// Validation done:
// If currently looked at value is not a number -> Exit program
func summer(nums []string, index int) int {
	if index < 0 {
		return 0
	}

	num, err := strconv.Atoi(nums[index])
	
	if err != nil {
		fmt.Println("Invalid Input")
		os.Exit(1)
	}

	if num < 0 {
		return summer(nums, index - 1)
	} else {
		return num * num + summer(nums, index - 1)
	}
	
}

// Function that will grab the size of the test case and the set of numbers in the test case and then run the summer function
// Validation done: 
// Check if size of test case is > 0, or not a number -> Exit program
// If test size is 0, Check if anything was submitted for the testcase, if so -> Exit program
// Check if size of test case is same as the size of the submitted set of numbers, if not -> Exit program
func cases(numOfCases int, answers []int) []int {
	if numOfCases == 0 {
		return answers
	}
	
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	numOfNums, err := strconv.Atoi(text)

	if err != nil || numOfNums < 0 {
		fmt.Println("Invalid Input")
		os.Exit(1)
	}

	if numOfNums == 0{
		answers = append(answers,0)
		emptyCheck, _ := reader.ReadString('\n')
		emptyCheck = strings.Replace(emptyCheck, "\r\n", "", -1)

		if emptyCheck == ""{
			return cases(numOfCases - 1, answers)
		}

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


// Main Function: Will grab the number of test cases from the user and run the cases function
// Validation done:
// If the value grabbed from the user is not a number or not positive -> Exit program
func main () {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r\n", "", -1)
	numOfCases, err := strconv.Atoi(text)
	
	if err != nil || numOfCases < 0 {
		fmt.Println("Invalid Input")
		os.Exit(1)
	}

	var answers []int

	answers = cases(numOfCases, answers)
	
	index := 0

	printAnswers(answers, index)

	os.Exit(1)	
	
}