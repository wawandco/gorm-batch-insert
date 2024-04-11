package internal

var (
	TCases = []struct {
		RecordsToCreate int
		Contacts        Contacts
	}{
		{RecordsToCreate: 100},
		{RecordsToCreate: 1_000},
		{RecordsToCreate: 10_000},
		{RecordsToCreate: 100_000},
		{RecordsToCreate: 300_000},
		{RecordsToCreate: 500_000},
		{RecordsToCreate: 1_000_000},
	}
)
