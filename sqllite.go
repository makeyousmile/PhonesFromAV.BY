package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
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
type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Newdb() DB {
	db := DB{}
	con, err := sql.Open("sqlite3", "db/avtodvor.db")
	db.sql = con
	checkErr(err)
	return db
}

//func dbInsertPhones(db DB, phones []Phone) {
//	stmt, err := db.sql.Prepare("INSERT OR IGNORE INTO phones(number, creation_time) values(?,?)")
//	db.stmt = stmt
//	checkErr(err)
//
//	for _, phone := range phones {
//		res, err := stmt.Exec(phone.number, phone.date)
//		checkErr(err)
//		added, err := res.RowsAffected()
//		checkErr(err)
//		if added == 1 {
//			log.Print(phone.number)
//			log.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
//		}
//
//	}
//
//}

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
			log.Print("номер попал в базу - отправляем в канал для послыки смс")
			//если номер попал в базу - отправляем в канал для послыки смс
			//sms <- "375" + phone.number
		}

	}

}
func GetPhonesCount() string {
	var data string
	rows := db.sql.QueryRow("SELECT COUNT(*) FROM phones")
	err := rows.Scan(&data)
	checkErr(err)
	return data

}
func SetSMS(sms string) {
	res, err := db.sql.Exec("UPDATE params SET sms = $1;", sms)
	checkErr(err)
	log.Print(res.RowsAffected())
}
func GetSMS() string {
	var sms string

	res := db.sql.QueryRow("select sms from params")
	err := res.Scan(&sms)
	checkErr(err)

	return sms
}
func GetTodayPhonesCount() string {
	var data string
	rows := db.sql.QueryRow("SELECT COUNT(*) FROM phones WHERE creation_time  >= DATE('now') AND creation_time < DATE('now', '+1 day')")
	err := rows.Scan(&data)
	checkErr(err)
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
	var baned []string

	rows, err := db.sql.Query("select number from phones WHERE baned = 1")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

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

func Cleardb() {
	res, err := db.sql.Exec("delete from phones")
	checkErr(err)
	log.Print(res.RowsAffected())
}

func addMessageId(number string, message_id string) {
	res, err := db.sql.Exec("UPDATE phones SET message_id = $1, send_time = $3 WHERE number = $2 ;", message_id, time.Now(), number)
	checkErr(err)
	log.Print(res.RowsAffected())
}
func SetRegions(region string, city string) {
	res, err := db.sql.Exec("UPDATE params SET regions = $1,cities = $2 ;", region, city)
	checkErr(err)
	added, _ := res.RowsAffected()
	if added != 0 {
		log.Print("update params")
		cfg.region = region
		cfg.city = city
	}
}
func GetRegions() string {
	var regions string

	res := db.sql.QueryRow("select regions from params")
	err := res.Scan(&regions)
	checkErr(err)
	return regions
}
func GetCities() string {
	var cities string

	res := db.sql.QueryRow("select cities from params")
	err := res.Scan(&cities)
	checkErr(err)
	return cities
}
func GetCitiesJson() string {
	var jsondata string
	var data []string
	var cities []City

	res := db.sql.QueryRow("select cities from params")
	err := res.Scan(&jsondata)
	checkErr(err)
	err = json.Unmarshal([]byte(jsondata), &data)
	for _, id := range data {
		checkErr(err)

		rows, err := db.sql.Query("select name from city WHERE id = $1", id)
		if err != nil {
			panic(err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)

		for rows.Next() {
			var city City
			city.Id, err = strconv.Atoi(id)
			checkErr(err)
			err := rows.Scan(&city.Name)
			if err != nil {
				checkErr(err)
				continue
			}

			cities = append(cities, city)
		}

	}
	fulljson, err := json.Marshal(cities)
	checkErr(err)
	return string(fulljson)
}
func GetCitiesWithRegion(jsondata string) string {
	var data []string
	var cities []City
	err := json.Unmarshal([]byte(jsondata), &data)
	checkErr(err)
	for _, region := range data {

		rows, err := db.sql.Query("select id,name from city WHERE region = $1", region)
		if err != nil {
			panic(err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)

		for rows.Next() {
			var city City
			err := rows.Scan(&city.Id, &city.Name)
			if err != nil {
				checkErr(err)
				continue
			}
			log.Print("city")
			log.Print(city.Name)
			cities = append(cities, city)
		}
	}
	fulljson, err := json.Marshal(cities)
	checkErr(err)

	return string(fulljson)

}
