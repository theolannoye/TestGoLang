package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

/* Clean the terminal with a unix command */
func cleanScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	cleanScreen()
	fmt.Println("Welcome in the game !")

	/* Var declaration */
	var myNumberMinimal, myNumberMaximal, numberToFind, testNumber, tryNumber int
	tryNumber = 0 /* Try counter */

	/* The player defines his min and max number */
	fmt.Println("Enter the minimal number: ")
	fmt.Scan(&myNumberMinimal) /* Read the standard entry. Pass a pointer to the variable we want to modify */
	fmt.Println("Enter the maximal number : ")
	fmt.Scan(&myNumberMaximal)

	/* We initialize the number to find */
	numberToFind = chooseANumber(myNumberMinimal, myNumberMaximal)

	/* The game is here. */
	for { /* Equivalent of while loop in Golang */
		fmt.Println("Which number are you thinking about ?")
		fmt.Scan(&testNumber)
		if(verifyNumber(numberToFind, testNumber)){
			tryNumber++
			break /*If the number has been find, we stop the loop and continue */
		} else if(testNumber < numberToFind){
			fmt.Println("Try more !")
		} else{
			fmt.Println("Try less !")
		}
		tryNumber++
	}

	fmt.Println("\nYou won in",tryNumber,"try")
	fmt.Println("The number was", numberToFind)

	/* Just to try wait groups and go routines*/
	a := sync.WaitGroup{}
	a.Add(5) /* We want 5 "done" to end the programm */
	j := 5 /* Counter to decrement (for the "x seconds left") */

	for i := 0; i<5; i++ {
		/* This function is executed in the same time in 5 threads */
		go func() {
			fmt.Println("The programm will end in 5 seconds left")
			time.Sleep(5*time.Second)
			j = j-1
			a.Done()
		}()
	}

	a.Wait() /* Execute the following lines if the  add delta is ok. */
	cleanScreen()
}
