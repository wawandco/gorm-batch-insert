package internal

import "github.com/jackc/fake"

func GenerateDummyContacts(numberOfContacts int) Contacts {
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
