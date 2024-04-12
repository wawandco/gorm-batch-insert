package benchmark_test

import (
	"fmt"
	"github/andrewmkano/gorm-batch-insert/internal"
	"github/andrewmkano/gorm-batch-insert/internal/gorman"
	"testing"
)

func BenchmarkSaveContactsInBatches(b *testing.B) {
	gbm, err := gorman.NewGormBatchManager()
	if err != nil {
		b.Fatal(err)
	}

	gbm.Conn.Exec("TRUNCATE contacts;")

	internal.PrepareTestCases()

	b.ResetTimer()

	for _, tcase := range internal.TCases {
		b.Run(fmt.Sprintf("records_number_%d", tcase.RecordsToCreate), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := gbm.SaveContactsInBatches(tcase.Contacts)
				if err != nil {
					b.Fatal(err.Error())
				}
			}
		})
	}
}
