package records

type NoSQLDatabase interface {
	Ping()
	InsertDocument()
	SearchAllDocument()
	DeleteDocument()
	UpdateDocument()
	DeleteManyUsingFilter()
}

type SQLDatabase interface {
	Ping()
	CreateTable()
	InsertRow()
	SearchAll()
	DeleteRow()
	UpdateRow()
}
