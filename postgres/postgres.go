package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected!")
	return &Db{db}, nil
}

func ConnString(host string, port string, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}

type Quote struct {
	ID        int
	Quotation string
	Person    string
}

func (d *Db) GetRandomQuote() Quote {
	query, err := d.Prepare("select * from quotes order by random() limit 1")

	if err != nil {
		fmt.Println("GetRandomQuote Preparation Err:", err)
	}

	row := query.QueryRow()

	var quote Quote
	err = row.Scan(
		&quote.ID,
		&quote.Quotation,
		&quote.Person,
	)

	if err != nil {
		fmt.Println("Error scaning row:", err)
	}

	return quote
}

func (d *Db) GetQuotes() []Quote {
	query, err := d.Prepare("SELECT * FROM quotes")
	if err != nil {
		fmt.Println("GetQuotes Preparation Err:", err)
	}

	rows, err := query.Query()
	if err != nil {
		fmt.Println("GetQuotes Query Err:", err)
	}

	var quote Quote
	quotes := []Quote{}

	for rows.Next() {
		err = rows.Scan(
			&quote.ID,
			&quote.Quotation,
			&quote.Person,
		)

		if err != nil {
			fmt.Println("Error scaning row:", err)
		}

		quotes = append(quotes, quote)
	}

	return quotes
}
