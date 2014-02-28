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

func TestCompany(t *testing.T) {
	company := ""
	company = Company()
	if company == "" {
		t.Errorf("Expected company to return a company name, Got %s", company)
	}
}

func ExampleUserF() {
	fmt.Println("Full Name", FullName())
	fmt.Println("First Name", FirstName())
	fmt.Println("Last Name", LastName())
	fmt.Println("Username", Username())
}

func ExampleCompany() {
	fmt.Println("Company", Company())
}
