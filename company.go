package fake

import (
	"fmt"
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
)

// Company returns a random company name.
func Company() (company string) {
	company = fmt.Sprintf("%s %s", randomString(COMPANY_NAMES), randomString(companySuffixes))
	return
}


