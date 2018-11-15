package models

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	err := InsertVoter("1", "5")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestGetAllPk(t *testing.T) {
	publicKeys, err := GetAllPk()
	if err != nil {
		t.Errorf("%v", err)
	}
	for _, value := range publicKeys {
		fmt.Println(value.X, value.Y)
	}
}
