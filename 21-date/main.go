package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Println(date)
	date = date.AddDate(1, 0, 0)
	fmt.Println(date)
}
