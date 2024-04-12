package internal

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
