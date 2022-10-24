package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	userAccount := CheckingAccount{
		"JP",
		123,
		321,
		10000,
	}

	fmt.Println(userAccount)
}
