# GORM Batch Insert Example 🧪

This is an alternative to the approach explained in the following post: [https://wawand.co/blog/posts/go-multi-inserting-in-batches/](https://wawand.co/blog/posts/go-multi-inserting-in-batches/)

Instead of using the sql library provided by Go, this time we're using an ORM called Gorm. Gorm v2.0 was announced a little while ago and one of the things that was mentioned was the [Batch insert support](https://gorm.io/docs/v2_release_note.html#Batch-Insert)

Since it has support for that by default, I decided to give it a go, following the same drill we ran last time.

## Requirements 🔎️

To run this project on your local environment, you need to have the following things installed:

- Go (This project uses version 1.22)
- PostgresApp
- PSQL

## Getting Started ⚙️

You will need to create an `.env` file and add the followign env variables we're using for this project:

```env
DB_USER
DB_PASSWORD
DB_HOST
DB_PORT
DB_NAME
BATCH_SIZE
```

The first five are used to build the DB's Data Source Name (DSN). The `BATCH_SIZE` serves to configure the size of the group you want GORM to use when inserting in batches.

### Running the app 🚀

As soon as you're done with the step above, run the following command:

```
go run cmd/setup/main.go
```

This will create the role needed for this drill, as well as the database and table we'll be using for the exercise at hand.

Now, with the following command you can start the app and have it insert a number of `contacts` in bulk.

```
# It will ask you for the number of records you wish to create.
go run cmd/app/main.go
```

## Results 📋

The following times were gathered using `time.Since`.
Here are the ones we obtained using the first approach(`database/sql`):
| Approach | Records Submitted | Time (ms) |
| ----------------------------- | ----------------- | ------------ |
| Single Multi-insert statement | 100 | 4.98245ms |
| Single Multi-insert statement | 10000 | 234.610741ms |
| Single Multi-insert statement | 20000 | 491.543135ms |

Now, here are the ones obtained using GORM and a Batch size of 500:

| Approach                  | Records Submitted | Time (ms)    |
| ------------------------- | ----------------- | ------------ |
| Batch Inserting with GORM | 10,000            | 67.422667ms  |
| Batch Inserting with GORM | 20,000            | 110.224666ms |
| Batch Inserting with GORM | 100,000           | 471.939583ms |
| Batch Inserting with GORM | 1,000,000         | 2.769831667s |

## Next Steps? 🧐

It seems that there is yet another alternative to this. From what I've read here and there, pgx has support for a [COPY protocol](https://pkg.go.dev/github.com/jackc/pgx/v4#hdr-Copy_Protocol). I haven't looked at it deeply, but perhaps that can be our next target in this series.

Built with ❤️ by Wawandco
