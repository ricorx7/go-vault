package main

import "fmt"

// CheckError will check the error message and display if
// found one.
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
