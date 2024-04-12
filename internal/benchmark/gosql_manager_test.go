package benchmark_test

// To try this one out, comment the other benchmark and uncomment this one.
// The idea is for these too to hold the same name
// so when we generate result text file we are able to get
// a comparison between these two benchmarks using benchstat.
// func BenchmarkSaveContactsInBatches(b *testing.B) {
// 	gsqbm, err := gsqlman.NewGoSQLBatchManager()
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	_, err = gsqbm.Conn.Exec("TRUNCATE contacts;")
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	internal.PrepareTestCases()

// 	b.ResetTimer()

// 	for _, tcase := range internal.TCases {
// 		b.Run(fmt.Sprintf("records_number_%d", tcase.RecordsToCreate), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				err := gsqbm.SaveContactsInBatches(tcase.Contacts)
// 				if err != nil {
// 					b.Fatal(err.Error())
// 				}
// 			}
// 		})
// 	}
// }
