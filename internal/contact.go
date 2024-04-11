package internal

import "github.com/jackc/fake"

type Contact struct {
	FirstName string
	LastName  string
	Email     string
}

type Contacts []Contact

func (c Contacts) SplitInGroups(groupSize int) []Contacts {
	if len(c) <= groupSize {
		return []Contacts{c}
	}

	var contactGroups []Contacts
	for i := 0; i < len(c); i += groupSize {
		end := i + groupSize
		if end > len(c) {
			end = len(c)
		}

		contactGroups = append(contactGroups, c[i:end])
	}

	return contactGroups
}

func generateDummyContacts(numberOfContacts int) Contacts {
	var contacts Contacts
	for i := 0; i < numberOfContacts; i++ {
		contacts = append(contacts, Contact{
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.EmailAddress(),
		})
	}

	return contacts
}

// PrepareTestCases loads all contacts in memory for each scenario.
func PrepareTestCases() {
	for i := range TCases {
		TCases[i].Contacts = generateDummyContacts(TCases[i].RecordsToCreate)
	}
}
