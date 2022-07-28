# SQLITE

Installation [Link](https://github.com/mattn/go-sqlite3)
```
go get github.com/mattn/go-sqlite3
```

Extra steps required for M1 Mac [Link](https://github.com/mattn/go-sqlite3)


## FAQ
You can register go functions as driver extensions [Link](https://pkg.go.dev/github.com/mattn/go-sqlite3#hdr-Go_SQlite3_Extensions)


Migrations [Link](https://github.com/golang-migrate/migrate)

Limitations
- Can't scan into string if the type may be nil as this is not supported by the library.
To work around this the type instead of a string must be list of btyes.