package fake

import "fmt"

var (
	FIRST_NAMES = []string{"Addison", "Alex", "Alexis", "Ali", "Amari", "Angel",
		"Ariel", "Ashton", "Avery", "Bailey", "Cameron", "Campbell", "Carson",
		"Casey", "Chase", "Christian", "Dakota", "Devin", "Devon", "Devyn",
		"Dominique", "Drew", "Dylan", "Emerson", "Guadalupe", "Harley",
		"Hayden", "Hunter", "Jaden", "Jadyn", "Jaiden", "Jaidyn", "Jaime",
		"Jamie", "Jayden", "Jaylen", "Jaylin", "Jessie", "Jordan", "Jordyn",
		"Justice", "Kasey", "Kayden", "Kendall", "Kennedy", "Logan", "London",
		"Madison", "Micah", "Morgan", "Parker", "Payton", "Peyton", "Phoenix",
		"Quinn", "Reagan", "Reese", "Reilly", "Riley", "Rowan", "Ryan", "Rylee",
		"Sage", "Shannon", "Shea", "Sidney", "Skylar", "Skyler", "Taylor",
		"Teagan", "Tyler"}

	LAST_NAMES = []string{"Smith", "Johnson", "Williams", "Jones", "Brown",
		"Davis", "Miller", "Wilson", "Moore", "Taylor", "Anderson", "Thomas",
		"Jackson", "White", "Harris", "Martin", "Thompson", "Garcia", "Martinez",
		"Robinson", "Clark", "Rodriguez", "Lewis", "Lee", "Walker", "Hall", "Allen",
		"Young", "Hernandez", "King", "Wright", "Lopez", "Hill", "Scott", "Green",
		"Adams", "Baker", "Gonzalez", "Nelson", "Carter", "Mitchell", "Perez",
		"Roberts", "Turner", "Phillips", "Campbell", "Parker", "Evans", "Edwards",
		"Collins", "Stewart", "Sanchez", "Morris", "Rogers", "Reed", "Cook", "Morgan",
		"Bell", "Murphy", "Bailey", "Rivera", "Cooper", "Richardson", "Cox", "Howard",
		"Ward", "Torres", "Peterson", "Gray", "Ramirez", "James", "Watson", "Brooks",
		"Kelly", "Sanders", "Price", "Bennett", "Wood", "Barnes", "Ross", "Henderson",
		"Coleman", "Jenkins", "Perry", "Powell", "Long", "Patterson", "Hughes",
		"Flores", "Washington", "Butler", "Simmons", "Foster", "Gonzales", "Bryant",
		"Alexander", "Russell", "Griffin", "Diaz", "Hayes", "Stevens"}

	usernameFormats = []string{"%s_%s", "%s.%s", "%s%s"}
)

// FullName returns a string formed from a random FIRST_NAME and LAST_NAME
// element.
func FullName() (name string) {
	name = FirstName() + " " + LastName()
	return
}

// FirstName returns a random string from the FIRST_NAME slice.
func FirstName() (name string) {
	name = randomString(FIRST_NAMES)
	return
}

// LastName returns a random string from the LAST_NAME slice.
func LastName() (name string) {
	name = randomString(LAST_NAMES)
	return
}

// Username random username.
func Username() (username string) {
	username = fmt.Sprintf(randomString(usernameFormats), randomString(FIRST_NAMES), randomString(LAST_NAMES))
	return
}

// Email generates a random email address for a user.
func Email() string {
	return fmt.Sprintf("%s@%s", Username(), Domain())
}
