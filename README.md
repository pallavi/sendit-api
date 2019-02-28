## climbing api

building this to learn go

## tools
you will need to install
- Go
- PostgreSQL

## first time setup
the first time you check out this project, install gometalinter and set up the database.
```sh
$ make tools 
$ make db-setup
```

to see all the makefile commands, run `make help`.

## running the server
install dependencies, update the database by migrating it to the latest version, and run the server.
```sh
$ make deps
$ make db-migrate
$ make
```

make requests to `localhost:3000`.

## migrations
to migrate the database, run
```sh
$ make db-migrate
```

to rollback the last set of migrations you ran, run
```sh
$ make db-rollback
```

to create a new migration (let's call it `add_new_table`), run
```sh
$ make db-migrate-create name=add_new_table
```
