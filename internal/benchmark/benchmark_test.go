package benchmark_test

import (
	"github/andrewmkano/gorm-batch-insert/internal"

	"github.com/jackc/fake"
)

var tcases = []struct {
	RecordsToCreate int
	Contacts        internal.Contacts
}{
	{RecordsToCreate: 100},
	{RecordsToCreate: 1_000},
	{RecordsToCreate: 10_000},
	{RecordsToCreate: 100_000},
	{RecordsToCreate: 300_000},
	{RecordsToCreate: 500_000},
	{RecordsToCreate: 1_000_000},
}

func generateDummyContacts(numberOfContacts int) internal.Contacts {
	var contacts internal.Contacts
	for i := 0; i < numberOfContacts; i++ {
		contacts = append(contacts, internal.Contact{
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.EmailAddress(),
		})
	}

	return contacts
}

func prepareTestCases() {
	for i := range tcases {
		tcases[i].Contacts = generateDummyContacts(tcases[i].RecordsToCreate)
	}
}
