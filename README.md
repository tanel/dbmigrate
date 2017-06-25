[![Build Status](https://travis-ci.org/tanel/dbmigrate.svg?branch=master)](https://travis-ci.org/tanel/dbmigrate)

There is a much better tool available now, check out [github.com/wallester/migrate](https://github.com/wallester/migrate)

Supported databases
-------------------
* PostgreSQL
* Cassandra

Install
-------
In your project, place your migrations in a separate folder,
for example, db/migrate.
**Migrations are sorted using their file name and then applied in the sorted order.**
Since sorting is important, name your migrations accordingly. For example,
add a timestamp before migration name. Or use any other ordering scheme you'll like.

Note that migration file names are saved into a table, and the table is used
later on to detect which migrations have already been applied. In other words,
**don't rename your migration files once they've been applied to your DB**.

Use
---

In your app code, import dbmigrate package:
```golang
import (
  "log"
  "path/filepath"

  "github.com/tanel/dbmigrate"
)
```

Then, run the migrations, depending on your database type.

Use with PostgreSQL
-------------------
**Make sure the migrations have an .sql ending.**

After app startup and after a sql.DB instance is initialized in your app, 
run the migrations. Assuming you have a variable called **db** that points to sql.DB
and the migrations are located in **db/migrate**, execute the following code:

```golang
if err := dbmigrate.Run(db, filepath.Join("db", "migrate")); err != nil {
  log.Fatal(err)
}
```

Use with Cassandra
------------------
**Make sure the migrations have an .cql ending.**

After app startup, open a session and run migrations:

```golang
session, err := cluster.CreateSession()
if err != nil {
	return err
}
defer session.Close()

cassandraMigrations := dbmigrate.NewCassandraDatabase(session, session)
if err := dbmigrate.ApplyMigrations(cassandraMigrations, filepath.Join("db", "migrate")); err != nil {
  return err
}
```
