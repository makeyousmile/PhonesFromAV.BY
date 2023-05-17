package main

import (
	"database/sql"
	"fmt"
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

func dbInsertPhone(phones chan Phone, sms chan string) {
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
			//sms <- "375296668485"
		}

	}

}
func GetPhonesCount() string {
	var data string
	rows := db.sql.QueryRow("SELECT COUNT(*) FROM phones")
	rows.Scan(&data)
	return data

}
func SetSMS(sms string) {
	res, err := db.sql.Exec("UPDATE params SET sms = $1;", sms)
	checkErr(err)
	log.Print(res.RowsAffected())
}
func GetSMS() string {
	var sms = ""

	res := db.sql.QueryRow("select sms from params")
	err := res.Scan(&sms)
	checkErr(err)

	return sms
}
func GetTodayPhonesCount() string {
	var data string
	rows := db.sql.QueryRow("SELECT COUNT(*) FROM phones WHERE creation_time  >= DATE('now') AND creation_time < DATE('now', '+1 day')")
	rows.Scan(&data)
	return data
}

func SetBaned(number string) {
	res, err := db.sql.Exec("UPDATE phones SET baned = 1 WHERE number = $1 ;", number)
	checkErr(err)
	found, _ := res.RowsAffected()
	if found == 0 {
		stmt, err := db.sql.Prepare("INSERT OR IGNORE INTO phones(number, creation_time, baned) values(?,?,?)")
		checkErr(err)
		_, err = stmt.Exec(number, time.Now(), 1)
		checkErr(err)

	}
}
func GetBaned() []string {
	baned := []string{}

	rows, err := db.sql.Query("select number from phones WHERE baned = 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		phone := ""
		err := rows.Scan(&phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		baned = append(baned, phone)
	}

	return baned
}
