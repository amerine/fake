package fake

import (
	"fmt"
	"testing"
)

func TestFullName(t *testing.T) {
	firstFullName := FullName()
	secondFullName := FullName()
	if firstFullName == secondFullName {
		t.Errorf("Expected FullName calls to be unique. Got %s and %s", firstFullName, secondFullName)
	}
}

func TestUserName(t *testing.T) {
	userName := Username()
	if userName == "" {
		t.Errorf("Expected UserName to return a user name. Got %s", userName)
	}
}

func ExampleUserF() {
	fmt.Println("Full Name", FullName())
	fmt.Println("First Name", FirstName())
	fmt.Println("Last Name", LastName())
	fmt.Println("Username", Username())
}
