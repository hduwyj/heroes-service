package models

import (
	"testing"
)

func TestUser(t *testing.T) {
	err := InsertVoter("1", "5")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestGetAllPk(t *testing.T) {

	UpdateGenerateKey("wang")
}
