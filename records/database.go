package records

type Database interface {
	Ping()
	InsertDocument()
	SearchAllDocument()
	DeleteDocument()
	UpdateDocument()
	DeleteManyUsingFilter()
}
