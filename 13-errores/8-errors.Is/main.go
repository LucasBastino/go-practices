package main

import (
	"errors"
	"fmt"
)

func main() {
	InsFundsError := PaymentError{msg: "Insufficient funds"}
	InvCreditCardError := PaymentError{msg: "Invalid credit card"}
	user := "customer"
	err := InsFundsError

	if user == "admin" {
		if errors.Is(err, InsFundsError) {
			fmt.Println("The user doesn't have money")
		}
		if errors.Is(err, InvCreditCardError) {
			fmt.Println("The user has a invalid card")
		}
	}

	if user == "customer" {
		fmt.Println()
	}
}

type PaymentError struct {
	msg string
}

func (pe PaymentError) Error() string {
	return fmt.Sprintf("An error has occurred: %v", pe.msg)
}
