package models

import (
	"log"
)

func InsertVoter(username string, password string) error {
	stmtIns, err := db.Prepare("INSERT INTO voter VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(username, password)
	if err != nil {
		return err
	}
	return nil
}

func IsValidateVoter(username string, password string) bool {
	stmtOut, err := db.Prepare("SELECT password FROM voter WHERE username = ?")
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	pwd := ""
	err = stmtOut.QueryRow(username).Scan(&pwd)
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	if pwd == password {
		return true
	}
	return false
}

//
//func UpdateVoter(voteCount int ,name string) error{
//	stmtOut,err := db.Prepare("UPDATE voter SET voteCount = ? WHERE name=?")
//	defer stmtOut.Close()
//	if err!=nil{
//		log.Printf("%v",err)
//		return err
//	}
//	voteCount++
//	_, err = stmtOut.Exec(voteCount, name)
//	if err!=nil{
//		log.Printf("%v",err)
//		return err
//	}
//}
