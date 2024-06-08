# Golang: Conditions

![image](https://github.com/luiscoco/Golang-sample4-Conditions/assets/32194879/7ab57146-f94b-41bb-95dd-baa07ab543c8)

## 1. Source code

```go
package main

import (
	"fmt"
)

func main() {
	// Variable to hold the user input
	var number int

	// Prompt the user to enter a number
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// If-Else statements to check if the number is positive, negative, or zero
	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}
```

## 2. How to run the application

```
go run conditions.go
```
