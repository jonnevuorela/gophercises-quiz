package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	targetFile := "problems.csv"
	var problemNumber int

	file, err := OpenFile(targetFile)
	CheckErr(err)
	problems, _ := ReadFile(file)
	problem := problems[problemNumber]
	file.Close()
	fmt.Println(problem[0])
}

/*Check Err*/
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*PickProblem()*/
func PickProblem(problems [][]string) int {
	var problemNumber int
	var currentProblemNumber int
	if currentProblemNumber < len(problems) {
		problemNumber = currentProblemNumber
		currentProblemNumber++
		return problemNumber
	}
	return problemNumber

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
	problems = append(problems, []string{"Game over", "0"})
	return problems, nil
}

/*Open file*/
func OpenFile(target string) (*os.File, error) {
	file, err := os.Open(target)
	CheckErr(err)
	return file, nil
}
