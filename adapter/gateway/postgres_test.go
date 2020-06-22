// +build integration

package gateway

import (
	"database/sql"
	"fmt"
	"jiroshin/clean-architecture-golang/entity"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
)

var sqlDB *sql.DB

func dbHandlingWrapper(m *testing.M) int {
	var err error

	sqlDB, err = sql.Open("postgres", dbConnString())
	if err != nil {
		log.Panic("Failed to connect to Postgres.")
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Panicf("DB connection failed %v", err)
	}

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(dbHandlingWrapper(m))
}

func TestStoreBook(t *testing.T) {
	storage := NewPostgres(sqlDB)

	var cases = []struct {
		Desc          string
		SampleBooks   []*entity.Book
		BookNum       int
		ExpectedError error
	}{
		{
			"insert Zero book",
			[]*entity.Book{},
			0,
			nil,
		},
		{
			"insert One book",
			[]*entity.Book{SampleBook1()},
			1,
			nil,
		},
		{
			"insert Different Two books",
			[]*entity.Book{SampleBook1(), SampleBook2()},
			2,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Desc, func(t *testing.T) {
			for i := 0; i < len(c.SampleBooks); i++ {
				book := c.SampleBooks[i]
				err := storage.StoreBook(book)
				assert.NoError(t, err)
			}

			results, err := getBooks(storage.DB)
			assert.NoError(t, err)

			assert.EqualValues(t, len(results), c.BookNum)

			err = DeleteBooks(storage.DB)
			assert.NoError(t, err)
		})
	}
}

func getBooks(db *sql.DB) ([]*entity.Book, error) {
	sql := `
		SELECT title, author, overview 
		FROM book;
	`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	var books []*entity.Book
	for rows.Next() {
		b := &entity.Book{}
		err := rows.Scan(&b.Title, &b.Author, &b.Overview)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

// func insertWorkspace(storage DataStorage, workspaceIdentifier string) error {
// 	sql := `
// 		INSERT INTO workspace (id, identifier)
// 		VALUES ($1, $2);
// 	`
//
// 	_, err := storage.DB.Exec(sql, uuid.FromStringOrNil(workspaceIdentifier), workspaceIdentifier)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func SampleBook1() *entity.Book {
	title := "サンプルブック1"
	author := "sugita"
	overview := "良い本でした"

	return &entity.Book{
		Title:    title,
		Author:   &author,
		Overview: &overview,
	}
}

func SampleBook2() *entity.Book {
	title := "サンプルブック2"
	author := "shinjiro"
	overview := "面白い本でした"

	return &entity.Book{
		Title:    title,
		Author:   &author,
		Overview: &overview,
	}
}

func dbConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		5432,
		os.Getenv("DB_NAME"))
}

func DeleteBooks(db *sql.DB) error {
	sql := "DELETE FROM book;"

	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}
