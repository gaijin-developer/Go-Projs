package main

import (
	"bufio"
	"fmt"
	"os"
	// "golang.org/x/text/language/display"
)

func IntroduceProgram(){
	fmt.Println("Welcome to Frankie todos")
	fmt.Println("1. Add a Todo")
	fmt.Println("2. Edit a Todo")
	fmt.Println("3. Delete a Todo")
	fmt.Println("5. Exit a Todo")
}

func getUserInput() string {

var userInput string = " "
fmt.Scanln(&userInput)

return userInput
}

func getLongerInput() string {
result,err  := bufio.NewReader(os.Stdin).ReadString('\n')

if err != nil {
	fmt.Println("there was an error")
	return "sorry invalid input"
}
return result	
}

func outputAQuestion(displayText string){
fmt.Println(displayText)
}

func AddATodo(todo string){
outputAQuestion("What would you like to do?")
userTodo := getLongerInput()

fmt.Println(userTodo)
}

func decideActionToTake(choice string){
// todoToBeAdded := getUserInput()

// fmt.Println(todoToBeAdded)
AddATodo(choice)
}