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

## Next Steps? 🧐

It seems that there is yet another alternative to this. From what I've read here and there, pgx has support for a [COPY protocol](https://pkg.go.dev/github.com/jackc/pgx/v4#hdr-Copy_Protocol). I haven't looked at it deeply, but perhaps that can be our next target in this series.

Built with ❤️ by Wawandco
