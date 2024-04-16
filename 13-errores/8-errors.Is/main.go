package main

import (
	"errors"
	"fmt"
)

var InsFundsError = PaymentError{msg: "Insufficient funds"}
var InvCreditCardError = PaymentError{msg: "Invalid credit card"}

type PaymentError struct {
	msg string
}

func (pe PaymentError) Error() string {
	return fmt.Sprintf("An error has occurred: %v", pe.msg)
}

func main() {

	user, err := dropError()
	if err != nil {
		if user == "admin" {
			if errors.Is(err, InsFundsError) {
				fmt.Println("The user doesn't have money")
			}
			if errors.Is(err, InvCreditCardError) {
				fmt.Println("The user has a invalid card")
			}
		}

		if user == "customer" {
			fmt.Println(err.Error())
		}
	}
}

func dropError() (string, error) {
	return "admin", InsFundsError
}
