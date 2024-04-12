# GORM vs Go/SQL Batch Insert Benchmark 🧪

This is a benchmark comparison between the go/sql approach used in this post: [https://wawand.co/blog/posts/go-multi-inserting-in-batches/](https://wawand.co/blog/posts/go-multi-inserting-in-batches/)

Instead of using the sql library provided by Go, this time we're using an ORM called Gorm. Gorm v2.0 was announced a little while ago and one of the things that was mentioned was the [Batch insert support](https://gorm.io/docs/v2_release_note.html#Batch-Insert)

## Requirements 🔎️

To run the benchmark tests, you need to have the following things installed:

- Go (This project uses version 1.22)
- PostgresApp
- PSQL
- Go Benchstat

## Getting Started ⚙️

You will need to create an `.env` file and add the following env variables we're using for this project:

```env
DB_USER
DB_PASSWORD
DB_HOST
DB_PORT
DB_NAME
BATCH_SIZE
```

The first five are used to build the DB's Data Source Name (DSN) for GORM and the database URL for the go/sql package. The `BATCH_SIZE` serves to configure the size of the group you want to use when inserting in batches(this affects both approaches).

## Running the benchmarks 📋

Both tests have the same name, hence why one of them is commented at the moment(since you can't have tests with the same name under the same package). They're also in the same package. The reason for all of that is to have benchstat analyze the results for both approaches in the form of an A/B test. So we run one benchmark after the other, saving the results to a text file and then have benchstat compare the results for us.

To run the benchmark test, use the following command:

```sh
$ GOMAXPROCS=2 go test -bench=Batches -timeout 30m -count 6 -benchtime=20x ./internal/benchmark | tee results/gorm-bench.txt
```

> ⚠️ Don't forget to change the name of the output file when trying to run the other benchmark.

`GOMAXPROCS` will tell the test suite to utilize two CPU cores to perform the benchmarks. We're also passing a `timeout` of `30m` since we're inserting quite the amount of records and don't want our benchmark to be interrupted(It's `11m` by default).

With the `-count` flag, we're asking the benchmark to run each scenario six times, each will have 20 loops to be executed, which is specified through the `benchtime` flag.

The `tee` command will help us see the results of our benchmarks while saving the results to a text file, so we can pass it to `benchstat` for the comparison.

## Results

To see the comparison between the two approaches, we used `benchstat` in the following fashion:

```sh
$ benchstat results/gsql-bench.txt results/gorm-bench.txt
```

Both files can be found in the `results` folder but feel free to generate your own if you please. Running the benchstat command should give you results in a the similar shape:

|                       -                        | results/gsql-bench.txt | results/gorm-bench.txt |         Delta         |
| :--------------------------------------------: | ---------------------- | ---------------------- | :-------------------: |
|                       -                        | sec/op                 | sec/op vs base         |           -           |
|   SaveContactsInBatches/records_number_100-2   | 1438.0µ ± 21%          | 974.8µ ± 20%           | -32.22% (p=0.004 n=6) |
|  SaveContactsInBatches/records_number_1000-2   | 5.606m ± 30%           | 5.029m ± 16%           |    ~ (p=0.093 n=6)    |
|  SaveContactsInBatches/records_number_10000-2  | 73.15m ± 25%           | 58.86m ± 4%            |    ~ (p=0.065 n=6)    |
| SaveContactsInBatches/records_number_100000-2  | 545.5m ± 15%           | 309.6m ± 30%           | -43.25% (p=0.002 n=6) |
| SaveContactsInBatches/records_number_300000-2  | 1.615 ± 2%             | 1.014 ± 8%             | -37.26% (p=0.002 n=6) |
| SaveContactsInBatches/records_number_500000-2  | 2.591 ± 6%             | 1.790 ± 12%            | -30.91% (p=0.002 n=6) |
| SaveContactsInBatches/records_number_1000000-2 | 5.130 ± 6%             | 3.786 ± 4%             | -26.19% (p=0.002 n=6) |
|                    geomean                     | 183.1m                 | 129.6m                 |        -29.24%        |

## Next Steps? 🧐

It seems that there is yet another alternative to this. From what I've read here and there, pgx has support for a [COPY protocol](https://pkg.go.dev/github.com/jackc/pgx/v4#hdr-Copy_Protocol). I haven't looked at it deeply, but perhaps that can be our next target in this series.

Built with ❤️ by Wawandco
