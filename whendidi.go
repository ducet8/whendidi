package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func displayResults(resultList []string) (){
	for index, result := range resultList {
		fmt.Printf("%d\t%s\n", index, string(result))
	}
}

func findExecutions(inputCommand string) (code int, commandStr string) {
	command := "grep " + inputCommand + " " + getHistoryPath()
	commandOut, _ := exec.Command("/bin/bash", "-c", command).Output()
	commandStr = string(commandOut)
	if commandStr == "" {
		fmt.Printf("Not found in your history!\n")
		code = 1
	} else {
		code = 0
	}
	return code, commandStr
}

func getHistoryPath() (fileName string) {
	usr, userErr := user.Current()
	if userErr == nil {
		fileName = usr.HomeDir + "/.bash_history"
	}
	return fileName
}

func getCommand(choice int, resultList []string) (command string){
	return resultList[choice]
}

func runChoice(choice string) () {
	commandOut, err := exec.Command("/bin/bash", "-c", choice).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf(string(commandOut))
}

func main() {
	oldCmd := os.Args[1]
	code, results := findExecutions(oldCmd)
	if code == 0 {
		resultList := strings.Split(strings.TrimSuffix(results, "\n"), "\n")
		displayResults(resultList)
		fmt.Print("\nWhich do you want to execute? > ")
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		choiceNum, _ := strconv.Atoi(strings.TrimSuffix(choice, "\n"))
		command := getCommand(choiceNum, resultList)
		runChoice(command)
	}
}