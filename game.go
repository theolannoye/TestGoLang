package main

import (
	"math/rand"
	"time"
)

/* Initialize a random number */
func chooseANumber(a, b int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(b-a) + a
}

/* Verify if 2 numbers are equals */
func verifyNumber(numberToFind int, numberChoosen int ) bool{
	return numberChoosen == numberToFind
}

