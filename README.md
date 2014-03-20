At the moment, stuff runs on Postgresql. About 2 lines of code would need be changed to support more DB-s. Pull requests ;)

Migration files
===============
Place your migration in a separate folder, for example, db/migrate.
Make sure the migrations have an .sql ending.
Migrations are sorted using their file name and then applied in the sorted order.
Since sorting is important, name your migrations accordingly. For example,
add a timestamp before migration name. Or use any other ordering scheme you'll like.

Run migrations
==============
In your app code, import dbmigrate package:
```golang
import (
  "github.com/tanel/dbmigrate"
  "log"
  "path/filepath"
)
```

Then, right after app startup and after a sql.DB instance is initialized in your app, 
run the migrations. Assuming you have a variable called **db** that points to sql.DB
and the migrations are located in **db/migrate**, execute the following code:

```golang
err := dbmigrate.Run(db, filepath.Join("db", "migrate")
if err != nil {
  log.Fatal(err)
}
```

Migrations are applied in a transactions. If any of these fail, the transaction is
rolled back. 

Also note that migration file names are saved into a table, and the table is used
later on to detect which migrations have already been applied. In other words,
**don't rename your migration files once they've been applied to your DB**.
