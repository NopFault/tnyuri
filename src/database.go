package tnyuri

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var config Config = GetConfig()

func instance() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", config.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func Init() {
	if _, err := os.Stat(config.Database); err != nil {
		file, err := os.Create(config.Database)
		if err != nil {
			log.Fatal("Something is wrong with Database file" + err.Error())
		}
		file.Close()
	}

	db := instance()

	createUrlTableSQL := `CREATE TABLE IF NOT EXISTS url (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url" TEXT,
		"short" TEXT NOT NULL UNIQUE,
		"user" TEXT,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createStatsTableSQL := `CREATE TABLE IF NOT EXISTS stats (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url_id" integer,
		"counter" integer
	);`

	urlStatement, err := db.Prepare(createUrlTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	urlStatement.Exec()

	statsStatement, err := db.Prepare(createStatsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statsStatement.Exec()

	defer db.Close()
}

func Insert(q string) int {
	db := instance()
	result, err := db.Exec(q)
	defer db.Close()

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Getting last Id (insert query): ", err)
	}

	return int(id)
}

func Exec(q string) {
	db := instance()
	_, err := db.Exec(q)
	defer db.Close()

	if err != nil {
		fmt.Println("Having problems by executing query:")
		fmt.Println(q)
		fmt.Println(err)
	}
}

func Select[R any](q string) R {
	var selection R
	db := instance()

	if err := db.Get(&selection, q); err != nil {
		fmt.Println("Error selecting data: ", err)
	}

	defer db.Close()

	return selection
}

func RowsBy[R any](from string, by string, val string) []R {
	db := instance()
	defer db.Close()

	var rows []R
	err := db.Select(&rows, "SELECT * FROM "+from+" where "+by+"='"+val+"'")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	return rows
}
