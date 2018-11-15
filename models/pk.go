package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"math/big"
)

var DefaultCurve = elliptic.P256()

func InsertPk(publicKeyX, publicKeyY string) error {
	stmtIns, err := db.Prepare("INSERT INTO pkey(publicKey_X,publicKey_Y) VALUES(?,?)")
	defer stmtIns.Close()
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(publicKeyX, publicKeyY)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPk() ([]ecdsa.PublicKey, error) {
	stmtOut, err := db.Prepare("SELECT publicKey_X,publicKey_Y FROM pkey")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	defer rows.Close()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	var XStr, YStr string
	var pk []ecdsa.PublicKey
	for rows.Next() {
		X := new(big.Int)
		Y := new(big.Int)
		err := rows.Scan(&XStr, &YStr)
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		X.SetString(XStr, 10)
		Y.SetString(YStr, 10)
		pk = append(pk, ecdsa.PublicKey{DefaultCurve, X, Y})
	}
	return pk, nil
}
