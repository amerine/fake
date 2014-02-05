package fake

import "testing"

func TestFullName(t *testing.T) {
	firstFullName := FullName()
	secondFullName := FullName()
	if firstFullName == secondFullName {
		t.Errorf("Expected FullName calls to be unique. Got %s and %s", firstFullName, secondFullName)
	}

}
