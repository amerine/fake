package fake

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	COMPANY_NAMES = []string{"CHOAM", "Acme", "Sirius Cybernetics", "Mom", "Rich",
		"Soylent", "Very Big Corp. of America", "Frozbozz Magic", "Warbucks", "Tyrell",
		"Wayne", "Virtucon", "Globex", "Umbrella", "Wonka", "Stark", "Clampett",
		"Oceanic Airlines", "Cyberdyne Systems", "Gringotts", "Oscorp", "Nakatomi Trading",
		"Spacely Space Sprockets",
	}

	companySuffixes = []string{"Corp.", "Industries", "Enterprises", "Inc", "LLC",
		"and Sons", "Group", "Ltd", "GmbH",
	}

	domainTLDs = []string{"net", "org", "com", "io"}
)

// Company returns a random company name.
func Company() (company string) {
	company = fmt.Sprintf("%s %s", randomString(COMPANY_NAMES), randomString(companySuffixes))
	return
}


// Domain returns a random domain name composed of a random company name, and
// a random TLD.
func Domain() (company string) {
	de := regexp.MustCompile("\\s|\\.")
	company = fmt.Sprintf("%s.%s", 
		strings.ToLower(de.ReplaceAllLiteralString(Company(), "")),
		randomTld())
	return
}

func randomTld() string {
	return randomString(domainTLDs)
}
