package models

import (
	"github.com/chainHero/heroes-service/web/controllers/util"
	"log"
)

func Query() []util.Candidate {

	//查询
	stmt, err := db.Prepare("select * from candidate ")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var candidates []util.Candidate
	candidate := util.Candidate{}
	for rows.Next() {
		err := rows.Scan(&candidate.Id, &candidate.Name, &candidate.Gender, &candidate.IdCard, &candidate.Content, &candidate.VoteCount)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(candidate)
		candidates = append(candidates, candidate)
	}
	return candidates
}

func InsertCandidate(c util.Candidate) error {
	stmt, err := db.Prepare("INSERT INTO candidate VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	_, err = stmt.Exec(c.Id, c.Name, c.Gender, c.IdCard, c.Content, c.VoteCount)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}
func DeleteCandidate(name string) error {
	stmt, err := db.Prepare("DELETE from candidate WHERE name = ?")
	defer stmt.Close()
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	_, err = stmt.Exec(name)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}

func UpdateCandidate(voteCount int, name string) error {
	stmt, err := db.Prepare("UPDATE candidate SET voteCount=? WHERE name = ?")
	defer stmt.Close()
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	voteCount++
	_, err = stmt.Exec(voteCount, name)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil
}
