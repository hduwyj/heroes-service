package controllers

import (
	"database/sql"
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers/util"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)
	cfg, err := ini.Load("conf/app.ini")
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	db, err = sql.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	Insert()
	Query()
}

func Query() {
	//查询
	stmt, err := db.Prepare("select * from candidate ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	candidate := util.Candidate{}
	for rows.Next() {
		err := rows.Scan(&candidate.Id, &candidate.Name, &candidate.Gender, &candidate.IdCard, &candidate.Content, &candidate.VoteCount)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(candidate)
	}

}

func Insert() {
	stmt, err := db.Prepare("INSERT INTO candidate VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(2, "特朗普", "男", "123456789", "请投特朗普一票", 0)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
