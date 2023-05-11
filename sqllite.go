package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type Phone struct {
	number string
	date   time.Time
}
type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
	buffer []Phone
}

func Newdb() DB {
	db := DB{}
	sql, err := sql.Open("sqlite3", "db/avtodvor.db")
	db.sql = sql
	checkErr(err)
	return db
}

func dbInsertPhones(db DB, phones []Phone) {
	stmt, err := db.sql.Prepare("INSERT OR IGNORE INTO phones(number, creation_time) values(?,?)")
	db.stmt = stmt
	checkErr(err)

	for _, phone := range phones {
		res, err := stmt.Exec(phone.number, phone.date)
		checkErr(err)
		added, err := res.RowsAffected()
		checkErr(err)
		if added == 1 {
			log.Print(phone.number)
			log.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		}

	}

}

func dbInsertPhone(phones chan Phone) {
	stmt, err := db.sql.Prepare("INSERT OR IGNORE INTO phones(number, creation_time) values(?,?)")
	db.stmt = stmt
	checkErr(err)

	for phone := range phones {
		res, err := stmt.Exec(phone.number, phone.date)
		checkErr(err)
		added, err := res.RowsAffected()
		checkErr(err)
		if added == 1 {
			log.Print(phone.number)
			log.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		}

	}

}
func GetPhonesCount() {

}
