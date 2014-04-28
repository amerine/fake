package fake

import (
	"fmt"
	"regexp"
	"strings"
)

var domainTLDs = []string{"net", "org", "com", "io"}

// Domain returns a random domain name composed of a random company name, and
// a random TLD.
func Domain() (domain string) {
	de := regexp.MustCompile("\\s|\\.")
	domain = fmt.Sprintf("%s.%s",
		strings.ToLower(de.ReplaceAllLiteralString(Company(), "")),
		randomTld())
	return
}

func randomTld() string {
	return randomString(domainTLDs)
}
