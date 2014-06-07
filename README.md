[![Build Status](https://travis-ci.org/tanel/dbmigrate.svg?branch=master)](https://travis-ci.org/tanel/dbmigrate)

How to use
==========
Place your migrations in a separate folder, for example, db/migrate.
**Make sure the migrations have an .sql ending.**
**Migrations are sorted using their file name and then applied in the sorted order.**
Since sorting is important, name your migrations accordingly. For example,
add a timestamp before migration name. Or use any other ordering scheme you'll like.

In your app code, import dbmigrate package:
```golang
import (
  "log"
  "path/filepath"

  "github.com/tanel/dbmigrate"
)
```

Then, right after app startup and after a sql.DB instance is initialized in your app, 
run the migrations. Assuming you have a variable called **db** that points to sql.DB
and the migrations are located in **db/migrate**, execute the following code:

```golang
if err := dbmigrate.Run(db, filepath.Join("db", "migrate")); err != nil {
  log.Fatal(err)
}
```

Note that migration file names are saved into a table, and the table is used
later on to detect which migrations have already been applied. In other words,
**don't rename your migration files once they've been applied to your DB**.
