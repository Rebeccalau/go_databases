package sqlite_implementation

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_databases/records"
	"log"
)

type SQLConnection struct{}

func (s *SQLConnection) CreateTable() {
	client := s.connect()
	defer client.Close()

	createTable := `CREATE TABLE if not exists record (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"primary" VARCHAR,
		"secondary" VARCHAR,
		"code" VARCHAR
	  );`

	stmt, err := client.Prepare(createTable)

	if err != nil {
		log.Fatalln(err)
	}

	result, err := stmt.Exec()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Successful Table Creation %s\n", result)
}

func (s *SQLConnection) InsertRow() {
	client := s.connect()
	defer client.Close()

	//insertRow := `INSERT INTO record ("primary", "code") VALUES ("1234-12324", "Code-1234");`
	insertRow := `INSERT INTO record ("primary", "secondary", "code") VALUES ("111-1111", "2222", "Code-3333" );`

	stmt, err := client.Prepare(insertRow)

	if err != nil {
		log.Fatalln(err)
	}

	result, err := stmt.Exec()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Successful Insert Row %s\n", result)
}

type SQLRow struct {
	id        int
	Primary   string
	Secondary sql.NullString // Or has to be []btye
	Code      string
}

func (s *SQLConnection) SearchAll() {
	client := s.connect()
	defer client.Close()

	selectQuery := `SELECT * FROM record;`

	stmt, err := client.Prepare(selectQuery)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err := stmt.Query()

	if err != nil {
		log.Fatalln(err)
	}

	result := []SQLRow{}
	for rows.Next() {
		var r SQLRow
		err = rows.Scan(&r.id, &r.Primary, &r.Secondary, &r.Code)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		result = append(result, r)
	}

	fmt.Println(result)
}

func (s *SQLConnection) DeleteRow() {
	client := s.connect()
	defer client.Close()

	insertRow := `DELETE FROM record WHERE "primary"="111-1111";`

	stmt, err := client.Prepare(insertRow)

	if err != nil {
		log.Fatalln(err)
	}

	result, err := stmt.Exec()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Successful Delete%s\n", result)
}

func (s *SQLConnection) UpdateRow() {
	client := s.connect()
	defer client.Close()

	insertRow := `UPDATE record set "code"="Code-update-1234" WHERE "primary"="1234-12324";`

	stmt, err := client.Prepare(insertRow)

	if err != nil {
		log.Fatalln(err)
	}

	result, err := stmt.Exec()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Successful Update%s\n", result)
}

func (s *SQLConnection) connect() *sql.DB {
	client, err := sql.Open("sqlite3", "./records.db")

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (s *SQLConnection) Ping() {
	client := s.connect()
	defer client.Close()

	err := client.Ping()

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successful Connection")
}

func NewSqlConnection() records.SQLDatabase {
	return &SQLConnection{}
}
