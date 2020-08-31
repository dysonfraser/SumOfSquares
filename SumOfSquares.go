package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main () {

	reader := bufio.NewReader(os.Stdin)
	
	numOfCases := 0

	for {

		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)
		numOfCases, _ = strconv.Atoi(text)

		var answers []int
		
		
		for i := 0; i < numOfCases; i++ {
			
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			
			numOfNums, _ := strconv.Atoi(text)

			muted, _ := reader.ReadString('\n')
			
			muted = strings.Replace(muted, "\r\n", "", -1)
			
			nums := strings.Split(muted, " ")
			
			sum := 0
			curNum := 0

			for j := 0; j < numOfNums; j++ {
				curNum, _ = strconv.Atoi(nums[j])
				if curNum > 0 {
					sum = sum + (curNum * curNum)		
				}
				
			}

			answers = append(answers, sum)
			
		}
		
		for i := 0; i < len(answers); i ++ {
			fmt.Println(answers[i])
		}

		os.Exit(1)
		
	}
}