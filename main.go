package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	targetFile := "problems.csv"
	currentProblemNumber := 0
	score := 0

	file, err := OpenFile(targetFile)
	CheckErr(err)
	problems, _ := ReadFile(file)
	file.Close()

	for i := 0; i < len(problems); i++ {
		AskQuestion(&currentProblemNumber, problems, &score)
	}

	fmt.Println("Game over. You've got " + strconv.FormatInt(int64(score), 10) + "/12 correct!")
}

/*Ask question*/
func AskQuestion(currentProblemNumber *int, problems [][]string, score *int) {
	currentProblem := PickProblem(problems, currentProblemNumber)
	fmt.Println("What is the correct answer?\n" + currentProblem[0])
	answer := GetInput()
	CheckAnswer(answer, currentProblem[1], score)
}

/*check answer*/
func CheckAnswer(answer int64, solution string, score *int) {
	sol, err := strconv.ParseInt(solution, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	if answer == sol {
		*score++
	}
}

/*ask input*/
func GetInput() int64 {
	var number int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your answer: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return number
	}
	input = strings.TrimSpace(input)
	number, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return number
}

/*Check Err*/
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*PickProblem()*/
func PickProblem(problems [][]string, currentProblemNumber *int) []string {
	var problem []string
	var problemNumber int
	if *currentProblemNumber < len(problems) {
		problemNumber = *currentProblemNumber
		*currentProblemNumber++
		problem = problems[problemNumber]
		return problem
	}
	return problem

}

/*ReadFile()*/
func ReadFile(file *os.File) ([][]string, error) {
	var problems [][]string
	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	CheckErr(err)
	for _, problem := range content {
		problems = append(problems, problem)
	}
	return problems, nil
}

/*Open file*/
func OpenFile(target string) (*os.File, error) {
	file, err := os.Open(target)
	CheckErr(err)
	return file, nil
}
