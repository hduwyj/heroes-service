package models

import (
	"log"
)

func InsertVoter(username string, password string) error {
	stmtIns, err := db.Prepare("INSERT INTO voter(username, password) VALUES(?,?)")
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
	defer stmtOut.Close()
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

func IsGenerateKey(username string) bool {
	stmtOut, err := db.Prepare("SELECT generateKey FROM voter WHERE username = ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%v", err)
		return false
	}
	generateKey := 0
	err = stmtOut.QueryRow(username).Scan(&generateKey)
	if generateKey == 0 {
		return false
	}
	return true
}

func UpdateGenerateKey(username string) error {
	stmtOut, err := db.Prepare("UPDATE voter SET generateKey = ? WHERE username=?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	_, err = stmtOut.Exec(1, username)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}
